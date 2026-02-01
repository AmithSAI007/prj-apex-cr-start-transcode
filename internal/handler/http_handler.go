package handler

import (
	"fmt"

	"go.uber.org/zap"

	"net/http"

	"github.com/AmithSAI007/prj-apex-cr-start-transcode.git/internal/config"
	"github.com/AmithSAI007/prj-apex-cr-start-transcode.git/internal/media"
	"github.com/AmithSAI007/prj-apex-cr-start-transcode.git/internal/storage"
	"github.com/AmithSAI007/prj-apex-cr-start-transcode.git/pkg/transcoder"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type StorageObjectData struct {
	Bucket string `json:"bucket"` // The GCS bucket that contains the object.
	Name   string `json:"name"`   // The name of the object.
}

type Handler struct {
	logger     *zap.Logger
	transcoder transcoder.TranscoderClient
	cfg        *config.Config
	client *storage.StorageClient
}

func NewHandler(logger *zap.Logger,
	transcoder transcoder.TranscoderClient,
	cfg *config.Config, client *storage.StorageClient) *Handler {
	return &Handler{
		logger:     logger,
		transcoder: transcoder,
		cfg:        cfg,
		client: client,
	}
}

func (h *Handler) HandleVideoProcessingEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "Expected HTTP POST request with CloudEvent payload", http.StatusMethodNotAllowed)
		return
	}

	event, err := cloudevents.NewEventFromHTTPRequest(r)
	if err != nil {
		h.logger.Error("Failed to parse CloudEvent", zap.Error(err))
		http.Error(w, "Failed to parse CloudEvent: "+err.Error(), http.StatusBadRequest)
		return
	}

	h.logger.Info("Received CloudEvent",
		zap.String("id", event.ID()),
		zap.String("source", event.Source()),
		zap.String("type", event.Type()),
	)

	var data StorageObjectData
	if err := event.DataAs(&data); err != nil {
		h.logger.Error("Failed to decode CloudEvent data", zap.Error(err))
		http.Error(w, "Failed to decode CloudEvent data: "+err.Error(), http.StatusBadRequest)
		return
	}


	httpUri, err := .GenerateSignedURL(data.Bucket, data.Name, 3600)
	if err != nil {
		h.logger.Error("Failed to generate signed URL", zap.Error(err))
		http.Error(w, "Failed to generate signed URL: "+err.Error(), http.StatusInternalServerError)
		return
	}


	inputURI := fmt.Sprintf("gs://%s/%s", data.Bucket, data.Name)

	hasAudio, err := media.HasAudioTrack(ctx, httpUri)
	if err != nil {
		h.logger.Error("Failed to analyze media file", zap.Error(err))
		http.Error(w, "Failed to analyze media file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !hasAudio {
		h.logger.Info("Media file has no audio track, skipping transcoding", zap.String("inputURI", inputURI))
		w.WriteHeader(http.StatusOK)
		return
	}

	jobName, err := h.transcoder.TriggerJobFromTemplate(
		ctx,
		h.cfg.ProjectID,
		h.cfg.Location,
		h.cfg.TemplateID,
		inputURI,
		h.cfg.OutputURI,
	)
	if err != nil {
		h.logger.Error("Failed to trigger transcoding job", zap.Error(err))
		http.Error(w, "Failed to trigger transcoding job: "+err.Error(), http.StatusInternalServerError)
		return
	}

	h.logger.Info("Transcoding job triggered successfully", zap.String("jobName", jobName))
	w.WriteHeader(http.StatusOK)

}
