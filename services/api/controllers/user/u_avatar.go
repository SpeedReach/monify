package user

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/lib"
	monify "monify/protobuf/gen/go"
	"monify/services/api/infra"
)

func (s Service) UpdateUserAvatar(ctx context.Context, req *monify.UpdateUserAvatarRequest) (*monify.UEmpty, error) {
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	imageService := ctx.Value(lib.FileServiceContextKey{}).(infra.FileService)
	_, err := imageService.ConfirmFileUsage(ctx, &monify.ConfirmFileUsageRequest{
		FileId: req.ImageId,
		Usage:  monify.Usage_UserAvatar,
		UserId: userId.String(),
	})
	if err != nil {
		return nil, err
	}
	return &monify.UEmpty{}, nil
}
