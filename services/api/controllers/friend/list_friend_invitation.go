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

func (s Service) ListFriendInvitation(ctx context.Context, req *monify.FriendEmpty) (*monify.ListFriendInvitationResponse, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	query, err := db.QueryContext(ctx, `
		SELECT nick_id, name, invite_id
		FROM user_identity JOIN friend_invite ON user_identity.user_id = friend_invite.sender
		WHERE friend_invite.receiver = $1`, userId)
	if err != nil {
		logger.Error("Select friend invitation error.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	defer query.Close()

	var invitaions []*monify.Invitation
	for {
		if !query.Next() {
			break
		}
		var invitaion monify.Invitation
		if err = query.Scan(&invitaion.SenderNickId, &invitaion.SenderName, &invitaion.InviteId); err != nil {
			logger.Error("Scan invitation nick_id and name error.", zap.Error(err))
			return nil, status.Error(codes.Internal, "")
		}
		invitaions = append(invitaions, &invitaion)
	}

	return &monify.ListFriendInvitationResponse{
		Invitation: invitaions,
	}, nil
}
