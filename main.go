package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AmithSAI007/prj-apex-cr-start-transcode.git/internal/config"
	"github.com/AmithSAI007/prj-apex-cr-start-transcode.git/internal/handler"
	"github.com/AmithSAI007/prj-apex-cr-start-transcode.git/pkg/transcoder"
	"go.uber.org/zap"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	logger, err := config.NewLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("Failed to sync logger: %v", err)
		}
	}()

	client, err := transcoder.NewTranscoderClient(context.Background(), logger)
	if err != nil {
		logger.Fatal("Failed to create Transcoder client", zap.Error(err))
	}
	defer func() {
		if err := client.Close(); err != nil {
			logger.Error("Failed to close Transcoder client", zap.Error(err))
		}
	}()

	handler := handler.NewHandler(logger, client, cfg)

	server := &http.Server{
		Addr:    ":" + cfg.HttpPort,
		Handler: http.HandlerFunc(handler.HandleVideoProcessingEvent),
	}

	go func() {
		logger.Info("Starting server", zap.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("server startup failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown signal received, starting graceful shutdown...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Fatal("server shutdown failed", zap.Error(err))
	}
	logger.Info("Server gracefully stopped")
}
