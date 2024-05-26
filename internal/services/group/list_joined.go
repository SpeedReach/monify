package group

import (
	"context"
	"database/sql"
	"monify/internal/middlewares"
	monify "monify/protobuf/gen/go"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) ListJoinedGroups(ctx context.Context, empty *monify.Empty) (*monify.ListJoinedGroupsResponse, error) {
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	userId := ctx.Value(middlewares.UserIdContextKey{})
	if userId == nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}

	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)
	query, err := db.Query(`
		SELECT "group".group_id, "group".name
		FROM "group" JOIN group_member ON "group".group_id = group_member.group_id
		WHERE user_id = $1`, userId)
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
		query.Scan(&group.GroupId, &group.Name)
		groups = append(groups, &group)
	}
	return &monify.ListJoinedGroupsResponse{
		Groups: groups,
	}, nil
}
