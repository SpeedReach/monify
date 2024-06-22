package infra

import (
	"context"
	"google.golang.org/grpc"
	"monify/lib/media"
	monify "monify/protobuf/gen/go"
)

type FileService struct {
	config Config
	client monify.MediaServiceClient
	media.FileHost
}

func (s FileService) ConfirmFileUsage(ctx context.Context, in *monify.ConfirmFileUsageRequest, opts ...grpc.CallOption) (*monify.MEmpty, error) {
	return s.client.ConfirmFileUsage(ctx, in, opts...)
}

func (s FileService) GetHost() string {
	return s.config.FileServerHost
}

func (s FileService) GetUrl(path string) string {
	return s.GetHost() + path
}
