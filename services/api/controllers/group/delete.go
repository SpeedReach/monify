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

func (s Service) DeleteGroup(ctx context.Context, req *monify.DeleteGroupRequest) (*monify.Empty, error) {
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid group id")
	}
	hasPerm, err := CheckPermission(ctx, groupId, userId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}
	if !hasPerm {
		return nil, status.Error(codes.PermissionDenied, "No permission")
	}

	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	result, err := db.ExecContext(ctx, `
	UPDATE "group" SET is_deleted = true WHERE group_id = $1 AND is_deleted = false`, groupId)
	if err != nil {
		logger.Error("Failed to delete group", zap.Error(err))
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("Failed to get rows affected", zap.Error(err))
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "Group not found")
	}

	return &monify.Empty{}, nil
}
