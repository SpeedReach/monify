package group_bill

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

func (s Service) ModifyGroupBill(ctx context.Context, req *monify.ModifyGroupBillRequest) (*monify.GroupGroupBillEmpty, error) {
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	/*
		userId, ok := ctx.Value(middlewares.UserIdContextKey{}).(uuid.UUID)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
		}
	*/
	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)

	/*
		//Check permission
		rows := db.QueryRowContext(ctx, `
			SELECT COUNT(*) FROM group_bill
			LEFT JOIN group_member gm ON group_bill.group_id = gm.group_id
			WHERE group_bill.bill_id = $1 AND gm.user_id = $2
		`, req.BillId, userId)
		var count int
		err := rows.Scan(&count)
		if err != nil {
			logger.Error("Failed to check permission", zap.Error(err))
		}
		if count != 1 {
			return nil, status.Error(codes.PermissionDenied, "No permission")
		}
	*/

	//START transaction
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		logger.Error("Failed to begin transaction", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	defer tx.Rollback()

	//Start modify (Delete -> Insert)
	//Delete
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

	_, err = tx.ExecContext(ctx,
		`DELETE FROM group_bill WHERE bill_id = $1`, req.BillId,
	)
	if err != nil {
		logger.Error("Failed to delete group from group_bill", zap.Error(err))
		return nil, err
	}

	//Insert
	bill_id, err := uuid.Parse(req.BillId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid bill id")
	}
	group_id, err := uuid.Parse(req.GroupId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid bill id")
	}
	member_id, err := uuid.Parse(req.MemberId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid bill id")
	}

	if err = insertBill(ctx, tx, logger, insertBillInfo{
		billId:        bill_id,
		groupId:       group_id,
		createdBy:     member_id,
		totalMoney:    req.TotalMoney,
		title:         req.Title,
		description:   req.Description,
		splitPeople:   req.SplitPeople,
		prepaidPeople: req.PrepaidPeople,
	}); err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	//Commit
	if err = tx.Commit(); err != nil {
		logger.Error("Failed to commit transaction", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}

	return &monify.GroupGroupBillEmpty{}, err
}
