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

func (s Service) ModifyGroupBill(ctx context.Context, req *monify.ModifyGroupBillRequest) (*monify.GroupGroupBillEmpty, error) {
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
	groupId, createdBy, err := getBillGroupIdAndCreator(ctx, billId)
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

	//Start modify (Delete -> Insert)
	//Delete
	var title string
	if title, err = deleteBill(ctx, tx, billId); err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	//Insert
	if err = insertBill(ctx, tx, insertBillInfo{
		billId:        billId,
		groupId:       groupId,
		createdBy:     createdBy,
		totalMoney:    req.TotalMoney,
		title:         req.Title,
		description:   req.Description,
		splitPeople:   req.SplitPeople,
		prepaidPeople: req.PrepaidPeople,
	}); err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	if err = processGroupBillModifyEvent(ctx, tx, group_bill.GroupBillModification{
		BillId:           billId,
		Ty:               group_bill.Update,
		OperatorMemberId: memberId,
		Title:            title,
		GroupId:          groupId,
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
