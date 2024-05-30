package group_bill

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/internal/middlewares"
	monify "monify/protobuf/gen/go"
)

func (s *Service) CreateBill(ctx context.Context, req *monify.CreateGroupBillRequest) (*monify.CreateGroupBillResponse, error) {
	userId, ok := ctx.Value(middlewares.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid group id")
	}
	if err = validateGroupBill(req); err != nil {
		return nil, err
	}

	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	db := ctx.Value(middlewares.StorageContextKey{}).(*sql.DB)
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		logger.Error("Failed to begin transaction", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	defer tx.Rollback()

	billId := uuid.New()
	_, err = tx.ExecContext(ctx, `
		INSERT INTO group_bill (bill_id, group_id, created_by, total_money, title, description)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, billId, groupId, userId, req.TotalMoney, req.Title, req.Description)
	if err != nil {
		logger.Error("Failed to insert group bill", zap.Error(err))
		return nil, err
	}
	for _, splitPerson := range req.SplitPeople {
		_, err = tx.ExecContext(ctx, `
			INSERT INTO group_split_bill (bill_id, person, amount) VALUES ($1, $2, $3)
		`, billId, splitPerson.MemberId, splitPerson.Amount)
		if err != nil {
			logger.Error("Failed to insert group split bill", zap.Error(err))
			return nil, err
		}
	}

	for _, prepaidPerson := range req.PrepaidPeople {
		_, err = tx.ExecContext(ctx, `
			INSERT INTO group_prepaid_bill (bill_id, person, amount) VALUES ($1, $2, $3)
			`, billId, prepaidPerson.MemberId, prepaidPerson.Amount)
		if err != nil {
			logger.Error("Failed to insert group prepaid bill", zap.Error(err))
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		logger.Error("Failed to commit transaction", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}

	return &monify.CreateGroupBillResponse{
		BillId: billId.String(),
	}, nil
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
