package group_bill

import (
	"context"
	"database/sql"
	"monify/internal/middlewares"
	monify "monify/protobuf/gen/go"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) DeleteBill(ctx context.Context, req *monify.DeleteGroupBillRequest) (*monify.GroupGroupBillEmpty, error) {
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)

	userId := ctx.Value(middlewares.UserIdContextKey{})
	if userId == nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}

	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		logger.Error("Failed to begin transaction", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	defer tx.Rollback()

	crearte_group_bill_response := &monify.CreateGroupBillResponse{}
	req.BillId = crearte_group_bill_response.GetBillId()

	_, err = tx.ExecContext(ctx,
		`DELETE FROM group_bill WHERE bill_id = $1`, req.BillId,
	)
	if err != nil {
		logger.Error("Failed to delete group from group_bill", zap.Error(err))
		return nil, err
	}

	_, err = tx.ExecContext(ctx,
		`DELETE FROM group_split_bill WHERE bill_id = $1`, req.BillId,
	)
	if err != nil {
		logger.Error("Failed to delete group from group_split_bill", zap.Error(err))
		return nil, err
	}

	_, err = tx.ExecContext(ctx,
		`DELETE FROM group_prepaid_bill WHERE bill_id = $1`, req.BillId,
	)
	if err != nil {
		logger.Error("Failed to delete group from group_prepaid_bill", zap.Error(err))
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		logger.Error("Failed to commit transaction", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}

	return &monify.GroupGroupBillEmpty{}, err
}
