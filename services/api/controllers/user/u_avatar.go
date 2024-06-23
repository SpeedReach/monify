package user

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go.uber.org/zap"
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
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	err := imageService.ConfirmFileUsage(ctx, &monify.ConfirmFileUsageRequest{
		FileId: req.ImageId,
		Usage:  monify.Usage_UserAvatar,
		UserId: userId.String(),
	})
	if err != nil {
		logger.Error("failed to confirm file usage", zap.Error(err))
		return nil, err
	}

	_, err = db.ExecContext(ctx, "UPDATE user_identity SET avatar_id = $1 WHERE user_id = $2", req.ImageId, userId.String())
	if err != nil {
		logger.Error("failed to update user avatar", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to update user avatar")
	}
	return &monify.UEmpty{}, nil
}
