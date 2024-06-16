package group_bill

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/lib"
	"monify/lib/group_bill"
	monify "monify/protobuf/gen/go"
	"monify/services/api/controllers/group"
)

// CreateGroupBill Handler
func (s Service) CreateGroupBill(ctx context.Context, req *monify.CreateGroupBillRequest) (*monify.CreateGroupBillResponse, error) {
	if req.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "Title is required")
	}
	//Parse ids
	userId, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid group id")
	}

	//Validation
	if err = validateGroupBill(req); err != nil {
		return nil, err
	}
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)

	//Check permission & get group_member_id of bill creator
	memberId, err := group.GetMemberId(ctx, groupId, userId)
	if err != nil {
		logger.Error("Failed to get member id", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	if memberId == uuid.Nil {
		return nil, status.Error(codes.PermissionDenied, "Permission denied")
	}

	//Begin Tx
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		logger.Error("Failed to begin transaction", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	defer tx.Rollback()
	//Insert
	billId := uuid.New()
	if err = insertBill(ctx, tx, insertBillInfo{
		billId:        billId,
		groupId:       groupId,
		createdBy:     memberId,
		totalMoney:    req.TotalMoney,
		title:         req.Title,
		description:   req.Description,
		splitPeople:   req.SplitPeople,
		prepaidPeople: req.PrepaidPeople,
	}); err != nil {
		return nil, status.Error(codes.Internal, "Internal")
	}

	if err = processGroupBillModifyEvent(ctx, tx, group_bill.GroupBillModification{
		GroupId:          groupId,
		Ty:               group_bill.Create,
		OperatorMemberId: memberId,
		BillId:           billId,
		Title:            req.Title,
	}); err != nil {
		logger.Error("", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}

	if err = tx.Commit(); err != nil {
		logger.Error("Failed to commit transaction", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}

	return &monify.CreateGroupBillResponse{
		BillId: billId.String(),
	}, nil
}

type insertBillInfo struct {
	billId        uuid.UUID
	groupId       uuid.UUID
	createdBy     uuid.UUID
	totalMoney    float64
	title         string
	description   string
	splitPeople   []*monify.InsertSplitPerson
	prepaidPeople []*monify.InsertPrepaidPerson
}

func insertBill(ctx context.Context, tx *sql.Tx, info insertBillInfo) error {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	_, err := tx.ExecContext(ctx, `
		INSERT INTO group_bill (bill_id, group_id, created_by, total_money, title, description)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, info.billId, info.groupId, info.createdBy, info.totalMoney, info.title, info.description)
	if err != nil {
		logger.Error("", zap.Error(err))
		return err
	}

	for _, splitPerson := range info.splitPeople {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO group_split_bill (bill_id, person, amount) VALUES ($1, $2, $3)
		`, info.billId, splitPerson.MemberId, splitPerson.Amount)
		if err != nil {
			logger.Error("Failed to insert group split bill", zap.Error(err))
			return err
		}
	}

	for _, prepaidPerson := range info.prepaidPeople {
		_, err = tx.ExecContext(ctx, `
			INSERT INTO group_prepaid_bill (bill_id, person, amount) VALUES ($1, $2, $3)
			`, info.billId, prepaidPerson.MemberId, prepaidPerson.Amount)
		if err != nil {
			logger.Error("Failed to insert group prepaid bill", zap.Error(err))
			return err
		}
	}

	return nil
}

func validateGroupBill(req *monify.CreateGroupBillRequest) error {
	if req.Title == "" {
		return status.Error(codes.InvalidArgument, "Title is required")
	}
	if req.TotalMoney <= 0 {
		return status.Error(codes.InvalidArgument, "Total money must be greater than 0")
	}
	if len(req.SplitPeople) == 0 {
		return status.Error(codes.InvalidArgument, "Split people is required")
	}
	if len(req.PrepaidPeople) == 0 {
		return status.Error(codes.InvalidArgument, "Prepaid people is required")
	}
	totalPrepaid := 0.0
	for _, prepaidPerson := range req.PrepaidPeople {
		totalPrepaid += prepaidPerson.Amount
	}
	if totalPrepaid != req.TotalMoney {
		return status.Error(codes.InvalidArgument, "Total money must be equal to total prepaid")
	}
	totalSplit := 0.0
	for _, splitPerson := range req.SplitPeople {
		totalSplit += splitPerson.Amount
	}
	if totalSplit != req.TotalMoney {
		return status.Error(codes.InvalidArgument, "Total money must be equal to total split")
	}
	return nil
}
