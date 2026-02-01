package storage

import (
	"context"
	"time"

	"cloud.google.com/go/storage"
	"go.uber.org/zap"
)

// StorageClient wraps the GCS client and centralizes signed URL generation.
type StorageClient struct {
	logger *zap.Logger
	client *storage.Client
}

// NewStorageClient initializes the Google Cloud Storage client.
func NewStorageClient(ctx context.Context, logger *zap.Logger) (*StorageClient, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		logger.Error("Failed to create GCS client", zap.Error(err))
		return nil, err
	}

	return &StorageClient{
		logger: logger,
		client: client,
	}, nil
}

// GenerateSignedURL creates a short-lived signed URL for a given object.
func (c *StorageClient) GenerateSignedURL(bucketName, objectName string, validDuration int64) (string, error) {
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(time.Duration(validDuration) * time.Minute),
	}

	c.logger.Info("Generating signed URL for object",
		zap.String("bucket", bucketName),
		zap.String("object", objectName),
		zap.Int64("validMinutes", validDuration),
	)

	url, err := c.client.Bucket(bucketName).SignedURL(objectName, opts)
	if err != nil {
		c.logger.Error("Failed to generate signed URL", zap.Error(err))
		return "", err
	}

	return url, nil
}

// Close releases the underlying GCS client resources.
func (c *StorageClient) Close() error {
	return c.client.Close()
}
