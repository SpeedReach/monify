package user

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/lib"
	monify "monify/protobuf/gen/go"
)

func (s Service) UpdateUserName(ctx context.Context, req *monify.UpdateUserNameRequest) (*monify.UEmpty, error) {
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}

	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	_, err := db.ExecContext(ctx, "UPDATE user_identity SET name = $1 WHERE user_id = $2", req.Name, userId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}
	return &monify.UEmpty{}, nil
}
