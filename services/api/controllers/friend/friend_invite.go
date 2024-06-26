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

func (s Service) InviteFriend(ctx context.Context, req *monify.InviteFriendRequest) (*monify.InviteFriendResponse, error) {

	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	userId := ctx.Value(lib.UserIdContextKey{})
	if userId == nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	receiverEmail := req.GetReceiverEmail()
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)

	query, err := db.QueryContext(ctx,
		`SELECT user_id FROM email_login WHERE email = $1`, receiverEmail)
	if err != nil {
		logger.Error("Select user_id by email error.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	query.Next()
	var receiverId string // uuid.UUIDs
	if err = query.Scan(&receiverId); err != nil {
		logger.Error("Scan receiver_id error.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	inviteId := uuid.New()
	_, err = db.ExecContext(ctx,
		`INSERT INTO friend_invite (invite_id, sender, receiver) VALUES ($1, $2, $3)`,
		inviteId, userId, receiverId)
	if err != nil {
		logger.Error("Failed to insert into friend_invite.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &monify.InviteFriendResponse{
		InviteId: inviteId.String(),
	}, nil
}
