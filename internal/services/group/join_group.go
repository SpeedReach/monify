package group

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/internal/middlewares"
	"monify/internal/utils"
	monify "monify/protobuf/gen/go"
	"time"
)

func (s Service) JoinGroup(ctx context.Context, req *monify.JoinGroupRequest) (*monify.JoinGroupResponse, error) {
	userId, ok := ctx.Value(middlewares.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	var groupId uuid.UUID
	var err error
	if groupId, err = findActiveInviteCode(ctx, db, logger, req.InviteCode); err != nil {
		return nil, err
	}
	var memberId uuid.UUID
	if memberId, err = createGroupMember(ctx, db, logger, groupId, userId); err != nil {
		return nil, err
	}
	return &monify.JoinGroupResponse{
		GroupId:  groupId.String(),
		MemberId: memberId.String(),
	}, nil
}

func findActiveInviteCode(ctx context.Context, db *sql.DB, logger *zap.Logger, inviteCode string) (uuid.UUID, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT group_id, created_at FROM group_invite_code WHERE invite_code = $1
	`, inviteCode)
	if err != nil {
		logger.Error("Failed to query group invite code", zap.Error(err))
		return uuid.Nil, status.Error(codes.Internal, "Internal")
	}
	defer rows.Close()
	if !rows.Next() {
		return uuid.Nil, status.Error(codes.NotFound, "Group invite code not found")
	}
	var groupId uuid.UUID
	var createdAt time.Time
	err = rows.Scan(&groupId, &createdAt)
	if err != nil {
		logger.Error("Failed to scan group invite code", zap.Error(err))
		return uuid.Nil, status.Error(codes.Internal, "Internal")
	}
	if time.Since(createdAt) > time.Duration(expiresInterval) {
		return uuid.Nil, status.Error(codes.InvalidArgument, "Group invite code expired")
	}

	return groupId, nil
}

func createGroupMember(ctx context.Context, db *sql.DB, logger *zap.Logger, groupId uuid.UUID, userId uuid.UUID) (uuid.UUID, error) {
	memberId := uuid.New()
	_, err := db.Exec(`
		INSERT INTO group_member (group_member_id, group_id, user_id) VALUES ($1,$2,$3)
	`, memberId, groupId, userId)
	if utils.IsDuplicateKeyError(err) {
		return uuid.Nil, status.Error(codes.AlreadyExists, "member already exists")
	}
	if err != nil {
		logger.Error("failed to insert ", zap.Error(err))
		return uuid.Nil, err
	}
	return memberId, nil
}
