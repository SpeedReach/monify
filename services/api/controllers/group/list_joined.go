package group

import (
	"context"
	"database/sql"
	"monify/lib"
	monify "monify/protobuf/gen/go"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) ListJoinedGroups(ctx context.Context, _ *monify.Empty) (*monify.ListJoinedGroupsResponse, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	userId := ctx.Value(lib.UserIdContextKey{})
	if userId == nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}

	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	query, err := db.QueryContext(ctx, `
		SELECT "group".group_id, "group".name, "group".description
		FROM "group" JOIN group_member ON "group".group_id = group_member.group_id
		WHERE user_id = $1 AND is_deleted = false`, userId)
	if err != nil {
		logger.Error("select group_id error", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	var groups []*monify.Group

	for {
		if !query.Next() {
			break
		}
		var group monify.Group
		if err = query.Scan(&group.GroupId, &group.Name, &group.Description); err != nil {
			logger.Error("scan group_id error", zap.Error(err))
			return nil, status.Error(codes.Internal, "")
		}
		groups = append(groups, &group)
	}
	return &monify.ListJoinedGroupsResponse{
		Groups: groups,
	}, nil
}
