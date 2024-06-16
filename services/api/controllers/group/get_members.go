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

func (s Service) GetGroupMembers(ctx context.Context, req *monify.GetGroupMembersRequest) (*monify.GetGroupMembersResponse, error) {
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Internal, "user id not found in context")
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid group id")
	}
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	permission, err := CheckPermission(ctx, groupId, userId)
	if err != nil {
		logger.Error("failed to check permission", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to check permission")
	}
	if !permission {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	rows, err := db.QueryContext(ctx, `
		SELECT gm.user_id, gm.group_member_id, ui.name
		FROM group_member gm
		LEFT JOIN user_identity ui on gm.user_id = ui.user_id
		WHERE group_id = $1
	`, groupId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to fetch group members")
	}
	defer rows.Close()
	var members []*monify.GroupMember
	for rows.Next() {
		var member monify.GroupMember
		if err := rows.Scan(&member.UserId, &member.MemberId, &member.UserName); err != nil {
			return nil, status.Error(codes.Internal, "failed to fetch group members")
		}
		members = append(members, &member)
	}
	return &monify.GetGroupMembersResponse{Members: members}, nil
}
