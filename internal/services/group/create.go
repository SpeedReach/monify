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

func createGroup(ctx context.Context, db *sql.Tx, name string) (uuid.UUID, error) {
	groupId := uuid.New()
	_, err := db.Exec(`
		INSERT INTO "group" (group_id, name) VALUES ($1, $2)
	`, groupId, name)
	return groupId, err
}

func (s Service) CreateGroup(ctx context.Context, req *monify.CreateGroupRequest) (*monify.CreateGroupResponse, error) {
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	userId, ok := ctx.Value(middlewares.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}

	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		logger.Error("", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()

	groupId, err := createGroup(ctx, tx, req.Name)
	if err != nil {
		logger.Error("", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal")
	}

	memberId, err := createGroupLeader(ctx, tx, groupId, userId)
	if err != nil {
		logger.Error("", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal")
	}

	if err = tx.Commit(); err != nil {
		logger.Error("failed to commit", zap.Error(err))
		return nil, err
	}

	return &monify.CreateGroupResponse{
		GroupId:  groupId.String(),
		MemberId: memberId.String(),
	}, nil
}

func createGroupLeader(ctx context.Context, db *sql.Tx, groupId uuid.UUID, userId uuid.UUID) (uuid.UUID, error) {
	groupMemberId := uuid.New()
	_, err := db.Exec(`
		INSERT INTO group_member (group_member_id,group_id, user_id) VALUES ($1,$2,$3)
	`, groupMemberId, groupId, userId)
	if err != nil {
		return uuid.Nil, err
	}
	return groupMemberId, nil
}
