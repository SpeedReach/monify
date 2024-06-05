package group_bill

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/internal/middlewares"
	"monify/internal/services/group"
	monify "monify/protobuf/gen/go"
)

func (s Service) GetHistory(ctx context.Context, req *monify.GetHistoryRequest) (*monify.GetHistoryResponse, error) {
	userId, ok := ctx.Value(middlewares.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid group id")
	}
	db := ctx.Value(middlewares.DatabaseContextKey{}).(*sql.DB)
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	havePerm, err := group.CheckPermission(ctx, groupId, userId)
	if err != nil {
		logger.Error("Failed to check permission", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	if !havePerm {
		return nil, status.Error(codes.PermissionDenied, "Permission denied")
	}

	rows, err := db.QueryContext(ctx, `
		SELECT type, title, ui.name, timestamp
		FROM group_bill_history
		LEFT JOIN group_member gm ON gm.group_member_id = group_bill_history.operator
		LEFT JOIN user_identity ui on gm.user_id = ui.user_id
		WHERE group_bill_history.group_id = $1
	`, groupId)
	if err != nil {
		logger.Error("Failed to select group bill history", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	defer rows.Close()

	var histories []*monify.GroupBillHistory
	for rows.Next() {
		var history monify.GroupBillHistory
		if err := rows.Scan(&history.Type, &history.Title, &history.Timestamp, &history.OperatorName); err != nil {
			logger.Error("Failed to scan group bill history", zap.Error(err))
			return nil, status.Error(codes.Internal, "Internal")
		}
		histories = append(histories, &history)
	}
	return &monify.GetHistoryResponse{Histories: histories}, nil
}
