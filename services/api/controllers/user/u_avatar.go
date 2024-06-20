package user

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/lib"
	monify "monify/protobuf/gen/go"
)

func (s Service) UpdateUserAvatar(ctx context.Context, req *monify.UpdateUserAvatarRequest) (*monify.UEmpty, error) {
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	imageService := ctx.Value(lib.ImageStorageContextKey{}).(monify.MediaServiceClient)
	_, err := imageService.ConfirmImageUsage(ctx, &monify.ConfirmImageUsageRequest{
		ImageId: req.ImageId,
		Usage:   monify.Usage_UserAvatar,
		UserId:  userId.String(),
	})
	if err != nil {
		return nil, err
	}

	//db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	return nil, nil
}
