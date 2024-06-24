package group_bill

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"monify/lib"
	monify "monify/protobuf/gen/go"
	"monify/services/api/controllers/group"
	"time"
)

func (s Service) GetHistory(ctx context.Context, req *monify.GetHistoryRequest) (*monify.GetHistoryResponse, error) {
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid group id")
	}
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
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
		var t time.Time
		if err := rows.Scan(&history.Type, &history.Title, &history.OperatorName, &t); err != nil {
			logger.Error("Failed to scan group bill history", zap.Error(err))
			return nil, status.Error(codes.Internal, "Internal")
		}
		history.Timestamp = timestamppb.New(t)
		histories = append(histories, &history)
	}
	return &monify.GetHistoryResponse{Histories: histories}, nil
}
