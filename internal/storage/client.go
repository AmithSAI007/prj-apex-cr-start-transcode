package storage

import (
	"context"
	"time"

	"cloud.google.com/go/storage"
	"go.uber.org/zap"
)

type StorageClient struct {
	logger *zap.Logger
	client *storage.Client
}

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

func (c *StorageClient) GenerateSignedURL(bucketName, objectName string, validDuration int64) (string, error) {

	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(time.Duration(validDuration) * time.Second),
	}

	url, err := c.client.Bucket(bucketName).SignedURL(objectName, opts)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (c *StorageClient) Close() error {
	return c.client.Close()
}
