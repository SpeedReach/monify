package friend

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

func (s Service) DeleteFriend(ctx context.Context, req *monify.DeleteFriendRequest) (*monify.FriendEmpty, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	_, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)

	_, err := db.ExecContext(ctx, `DELETE FROM friend WHERE relation_id = $1`, req.RelationId)
	if err != nil {
		logger.Error("Delete friend error.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &monify.FriendEmpty{}, nil

}
