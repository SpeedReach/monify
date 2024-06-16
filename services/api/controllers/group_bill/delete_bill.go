package group_bill

import (
	"context"
	"database/sql"
	"monify/lib"
	"monify/lib/group_bill"
	monify "monify/protobuf/gen/go"
	"monify/services/api/controllers/group"

	"github.com/google/uuid"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Service) DeleteGroupBill(ctx context.Context, req *monify.DeleteGroupBillRequest) (*monify.GroupGroupBillEmpty, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	billId, err := uuid.Parse(req.BillId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid bill id")
	}
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)

	//Check permission
	groupId, err := getBillGroupId(ctx, billId)
	if groupId == uuid.Nil {
		if err != nil {
			logger.Error("", zap.Error(err))
			return nil, status.Error(codes.Internal, "Internal")
		}
		return nil, status.Error(codes.NotFound, "Not found")
	}
	memberId, err := group.GetMemberId(ctx, groupId, userId)
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
	var title string
	if title, err = deleteBill(ctx, tx, billId); err != nil {
		logger.Error("", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}

	//insert history
	if err = processGroupBillModifyEvent(ctx, tx, group_bill.GroupBillModification{
		BillId:           billId,
		Ty:               group_bill.Delete,
		OperatorMemberId: memberId,
		Title:            title,
		GroupId:          groupId,
	}); err != nil {
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

// getBillGroupId returns the deleted title of the bill
func deleteBill(ctx context.Context, tx *sql.Tx, billId uuid.UUID) (string, error) {
	var title string
	if err := tx.QueryRowContext(ctx, "SELECT title FROM group_bill_history WHERE bill_id=$1", billId).Scan(&title); err != nil {
		return "", err
	}

	_, err := tx.ExecContext(ctx,
		`DELETE FROM group_split_bill WHERE bill_id = $1`, billId,
	)

	if err != nil {
		return "", err
	}

	_, err = tx.ExecContext(ctx,
		`DELETE FROM group_prepaid_bill WHERE bill_id = $1`, billId,
	)
	if err != nil {
		return "", err
	}

	_, err = tx.ExecContext(ctx,
		`DELETE FROM group_bill WHERE bill_id = $1`, billId,
	)

	if err != nil {
		return "", err
	}
	return title, err
}
