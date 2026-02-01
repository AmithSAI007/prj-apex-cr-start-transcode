# prj-apex-cr-start-transcode

## Overview
This service listens for Google Cloud Storage events delivered as CloudEvents
and triggers a Google Cloud Transcoder job when an uploaded video contains an
audio track. It skips transcoding for silent videos to save processing time.

## How it works
- Receives a CloudEvent HTTP POST from Cloud Storage.
- Generates a short-lived signed URL to inspect the media with `ffprobe`.
- If an audio stream exists, it starts a Transcoder job using a template.

## Configuration
Configuration is loaded from `config.yaml` and environment variables.

| Setting | Description |
| --- | --- |
| `APP_ENV` | Environment name (`development`, `production`, etc.). |
| `HTTP_PORT` | Port for the HTTP server (default `8080`). |
| `PROJECT_ID` | GCP project ID for Transcoder jobs. |
| `LOCATION` | GCP region for the Transcoder service. |
| `TEMPLATE_ID` | Transcoder template ID to apply. |
| `OUTPUT_URI` | GCS path for transcoder output. |

## Running locally
1. Install Go and ensure `ffprobe` is available in your PATH.
2. Update `config.yaml` or set environment variables.
3. Start the service:

```bash
go run .
```

The service listens on `HTTP_PORT` and expects CloudEvent POST requests.
