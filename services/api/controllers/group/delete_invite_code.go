package group

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

func (s Service) DeleteInviteCode(ctx context.Context, req *monify.DeleteInviteCodeRequest) (*monify.Empty, error) {
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

	if _, err = db.ExecContext(ctx, "DELETE FROM group_invite_code WHERE group_id = $1", groupId); err != nil {
		logger.Error("Failed to delete invite code", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}

	return &monify.Empty{}, nil
}
