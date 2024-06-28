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

func (s Service) RejectInvitation(ctx context.Context, req *monify.RejectInvitationRequest) (*monify.FriendEmpty, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	userId := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if userId == uuid.Nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)

	_, err := db.ExecContext(ctx, `DELETE FROM friend_invite WHERE invite_id = $1`, req.InviteId)
	if err != nil {
		logger.Error("Delete friend invitation error.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &monify.FriendEmpty{}, nil

}
