package user

import (
	"context"
	"database/sql"
	"monify/lib"
	monify "monify/protobuf/gen/go"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) UpdateUserNickId(ctx context.Context, req *monify.UpdateUserNickIdRequest) (*monify.UEmpty, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}

	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	_, err := db.ExecContext(ctx, "UPDATE user_identity SET nick_id = $1 WHERE user_id = $2", req.NickId, userId)
	if err != nil {
		logger.Error("Update nick_id error.", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	return &monify.UEmpty{}, nil
}
