package media

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type FFProbeOutput struct {
	Streams []Stream `json:"streams"`
}

type Stream struct {
	Index     int    `json:"index"`
	CodecType string `json:"codec_type"`
}

func HasAudioTrack(ctx context.Context, gcsURI string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "ffprobe",
		"-v", "error",
		"-show_entries", "stream=codec_type",
		"-of", "json",
		gcsURI,
	)

	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("ffprobe failed: %w", err)
	}

	var probe FFProbeOutput
	if err := json.Unmarshal(output, &probe); err != nil {
		return false, err
	}

	for _, s := range probe.Streams {
		if s.CodecType == "audio" {
			return true, nil
		}
	}

	return false, nil
}
