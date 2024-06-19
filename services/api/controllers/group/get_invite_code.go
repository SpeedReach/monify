package group

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"monify/lib"
	"monify/lib/group"
	monify "monify/protobuf/gen/go"
)

func (s Service) GetInviteCode(ctx context.Context, req *monify.GetInviteCodeRequest) (*monify.GetInviteCodeResponse, error) {
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid group ID")
	}
	hasPerm, err := CheckPermission(ctx, groupId, userId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}
	if !hasPerm {
		return nil, status.Error(codes.PermissionDenied, "Permission denied")
	}

	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)

	inviteCode := group.InviteCode{GroupId: groupId}
	err = db.QueryRowContext(ctx, "SELECT invite_code, created_at FROM group_invite_code WHERE group_id = $1", groupId).Scan(&inviteCode.Code, &inviteCode.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "Invite code not found")
		}
		logger.Error("Failed to get invite code", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}

	if inviteCode.IsExpired() {
		logger.Info("queried a expired invite code")
		return nil, status.Error(codes.NotFound, "Invite code expired")
	}

	return &monify.GetInviteCodeResponse{InviteCode: inviteCode.Code, ExpiresAfter: durationpb.New(inviteCode.ExpiresAfter())}, nil
}
