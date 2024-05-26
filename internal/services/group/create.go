package group

import (
	"context"
	"database/sql"
	"monify/internal/middlewares"
	monify "monify/protobuf/gen/go"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func createGroup(ctx context.Context, db *sql.DB, name string) (uuid.UUID, error) {
	groupId := uuid.New()
	_, err := db.Exec(`
		INSERT INTO "group" (group_id, name) VALUES ($1, $2)
	`, groupId, name)
	return groupId, err
}

func (s Service) CreateGroup(ctx context.Context, req *monify.CreateGroupRequest) (*monify.CreateGroupResponse, error) {
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	userId := ctx.Value(middlewares.UserIdContextKey{})
	if userId == nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}

	userUid := userId.(uuid.UUID)

	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)
	groupId, err := createGroup(ctx, db, req.Name)
	if err != nil {
		logger.Error("", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal")
	}

	memberId, err := createGroupMember(ctx, db, groupId, userUid)
	if err != nil {
		logger.Error("", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal")
	}

	return &monify.CreateGroupResponse{
		GroupId:  groupId.String(),
		MemberId: memberId.String(),
	}, nil
}
