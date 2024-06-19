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

func (s Service) GetGroupInfo(ctx context.Context, req *monify.GetGroupInfoRequest) (*monify.GetGroupInfoResponse, error) {
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

	return GetGroupInfo(ctx, groupId)
}

func (s Service) GetGroupByInviteCode(ctx context.Context, req *monify.GetGroupByInviteCodeRequest) (*monify.GetGroupInfoResponse, error) {
	_, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}

	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)

	var groupId uuid.UUID
	err := db.QueryRowContext(ctx, "SELECT group_id FROM group_invite_code WHERE invite_code = $1", req.InviteCode).Scan(&groupId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "Group not found")
		}
		logger.Error("Failed to get group by invite code", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}

	return GetGroupInfo(ctx, groupId)
}

func GetGroupInfo(ctx context.Context, groupId uuid.UUID) (*monify.GetGroupInfoResponse, error) {
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)

	response := &monify.GetGroupInfoResponse{
		GroupId: groupId.String(),
	}
	err := db.QueryRowContext(ctx, `SELECT name, description, avatar_url FROM "group" WHERE group_id = $1`, groupId).Scan(&response.Name, &response.Description, &response.AvatarUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "Group not found")
		}
		logger.Error("Failed to get group info", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}

	return response, nil
}
