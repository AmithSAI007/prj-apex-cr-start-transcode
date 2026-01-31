package transcoder

import (
	"context"

	transcoder "cloud.google.com/go/video/transcoder/apiv1"
	"cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
	"go.uber.org/zap"
)

type TranscoderClient interface {
	TriggerJobFromTemplate(ctx context.Context, projectID, location, templateID, inputURI, outputURI string) (string, error)
}

type Client struct {
	logger    *zap.Logger
	gcpClient *transcoder.Client
}

func NewTranscoderClient(ctx context.Context, logger *zap.Logger) (*Client, error) {
	client, err := transcoder.NewClient(ctx)
	if err != nil {
		logger.Error("Failed to create Transcoder client", zap.Error(err))
		return nil, err
	}

	return &Client{
		logger:    logger,
		gcpClient: client,
	}, nil
}

func (c *Client) Close() error {
	return c.gcpClient.Close()
}

var _ TranscoderClient = (*Client)(nil)

func (c *Client) TriggerJobFromTemplate(ctx context.Context, projectID, location, templateID, inputURI, outputURI string) (string, error) {

	c.logger.Info("Creating transcoding job from template",
		zap.String("projectID", projectID),
		zap.String("location", location),
		zap.String("templateID", templateID),
		zap.String("inputURI", inputURI),
		zap.String("outputURI", outputURI),
	)

	req := &transcoderpb.CreateJobRequest{
		Parent: "projects/" + projectID + "/locations/" + location,
		Job: &transcoderpb.Job{
			InputUri:  inputURI,
			OutputUri: outputURI,
			JobConfig: &transcoderpb.Job_TemplateId{
				TemplateId: templateID,
			},
		},
	}

	job, err := c.gcpClient.CreateJob(context.Background(), req)
	if err != nil {
		c.logger.Error("Failed to create transcoding job", zap.Error(err))
		return "", err
	}

	c.logger.Info("Transcoding job created successfully",
		zap.String("jobName", job.Name),
	)

	return job.Name, nil

}
