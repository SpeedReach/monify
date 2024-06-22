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
)

func (Service) GetUserInfo(ctx context.Context, req *monify.GetUserInfoRequest) (*monify.GetUserInfoResponse, error) {
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	response := monify.GetUserInfoResponse{}

	err := db.QueryRowContext(ctx, "SELECT name, avatar_url FROM user_identity WHERE user_id = $1", userId).Scan(&response.Name, &response.AvatarUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		logger.Error("failed to get user info", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	return &response, nil
}
