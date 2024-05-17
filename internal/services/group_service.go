package services

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/internal/middlewares"
	monify "monify/protobuf"
)

type GroupService struct {
	monify.UnimplementedGroupServiceServer
}

func (g GroupService) CreateGroup(ctx context.Context, req *monify.CreateGroupRequest) (*monify.CreateGroupResponse, error) {
	userId := ctx.Value(middlewares.UserIdContextKey{})
	if userId == nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}

	userUid := userId.(uuid.UUID)

	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)
	groupId := uuid.New()
	_, err := db.Exec(`
		INSERT INTO "group" (group_id, name) VALUES ($1, $2)
	`, groupId, req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	_, err = db.Exec(`
		INSERT INTO group_member (group_member_id, group_id, user_id) VALUES ($1,$2,$3)
	`, uuid.New(), groupId, userUid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	return &monify.CreateGroupResponse{
		GroupId: groupId.String(),
	}, nil
}
