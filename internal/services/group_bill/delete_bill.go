package group_bill

import (
	"context"
	"database/sql"
	"monify/internal/middlewares"
	"monify/internal/services/group"
	monify "monify/protobuf/gen/go"

	"github.com/google/uuid"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) DeleteGroupBill(ctx context.Context, req *monify.DeleteGroupBillRequest) (*monify.GroupGroupBillEmpty, error) {
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	userId, ok := ctx.Value(middlewares.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	billId, err := uuid.Parse(req.BillId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid bill id")
	}
	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)

	//Check permission
	groupId, err := getBillGroupId(ctx, db, billId)
	if groupId == uuid.Nil {
		if err != nil {
			logger.Error("", zap.Error(err))
			return nil, status.Error(codes.Internal, "Internal")
		}
		return nil, status.Error(codes.NotFound, "Not found")
	}
	memberId, err := group.GetMemberId(ctx, db, groupId, userId)
	if memberId == uuid.Nil {
		if err != nil {
			logger.Error("", zap.Error(err))
		}
		return nil, status.Error(codes.PermissionDenied, "No permission")
	}

	//START transaction
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		logger.Error("Failed to begin transaction", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	defer tx.Rollback()

	//Start delete
	if err = deleteBill(ctx, tx, billId); err != nil {
		logger.Error("", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}

	//Commit
	if err = tx.Commit(); err != nil {
		logger.Error("Failed to commit transaction", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}

	return &monify.GroupGroupBillEmpty{}, err
}

func deleteBill(ctx context.Context, tx *sql.Tx, billId uuid.UUID) error {
	_, err := tx.ExecContext(ctx,
		`DELETE FROM group_split_bill WHERE bill_id = $1`, billId,
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx,
		`DELETE FROM group_prepaid_bill WHERE bill_id = $1`, billId,
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx,
		`DELETE FROM group_bill WHERE bill_id = $1`, billId,
	)

	if err != nil {
		return err
	}
	return nil
}
