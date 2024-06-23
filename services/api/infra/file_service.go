package infra

import (
	"context"
	"monify/lib/media"
	monify "monify/protobuf/gen/go"
)

type FileService struct {
	config Config
	client monify.MediaServiceClient
	media.FileHost
}

func (s FileService) ConfirmFileUsage(ctx context.Context, in *monify.ConfirmFileUsageRequest) error {
	_, err := s.client.ConfirmFileUsage(ctx, in)
	return err
}

func (s FileService) GetHost() string {
	return s.config.FileServerHost
}

func (s FileService) GetUrl(path string) string {
	return s.GetHost() + "/" + path
}
