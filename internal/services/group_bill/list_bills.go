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

// GetGroupBills Handler
func (s Service) GetGroupBills(ctx context.Context, req *monify.GetGroupBillsRequest) (*monify.GetGroupBillsResponse, error) {
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

	permission, err := group.CheckPermission(ctx, groupId, userId)
	if err != nil {
		logger.Error("Failed to check permission", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	if !permission {
		return nil, status.Error(codes.PermissionDenied, "Permission denied")
	}

	bills, err := getGroupBills(ctx, db, logger, groupId)
	if err != nil {
		return nil, err
	}
	return &monify.GetGroupBillsResponse{GroupBills: bills}, nil
}

func getGroupBills(ctx context.Context, db *sql.DB, logger *zap.Logger, groupId uuid.UUID) ([]*monify.CreatedGroupBill, error) {
	var bills []*monify.CreatedGroupBill
	query, err := db.Query(`
		SELECT bill_id, total_money, title, description
		FROM group_bill
		WHERE group_id = $1
	`, groupId)
	if err != nil {
		logger.Error("Failed to select group bills", zap.Error(err))
		return bills, status.Error(codes.Internal, "Internal")
	}
	defer query.Close()

	for query.Next() {
		var billId uuid.UUID
		var totalMoney float64
		var title, description string
		err = query.Scan(&billId, &totalMoney, &title, &description)
		if err != nil {
			logger.Error("Failed to scan group bill", zap.Error(err))
			return nil, status.Error(codes.Internal, "Internal")
		}
		splits, err := getGroupBillSplits(ctx, db, logger, billId)
		if err != nil {
			return nil, err
		}
		println(len(splits))
		prepaid, err := getGroupBillPrepaid(ctx, db, logger, billId)
		if err != nil {
			return nil, err
		}
		bill := monify.CreatedGroupBill{
			BillId:        billId.String(),
			TotalMoney:    totalMoney,
			Title:         title,
			Description:   description,
			SplitPeople:   splits,
			PrepaidPeople: prepaid,
		}

		bills = append(bills, &bill)
	}
	return bills, nil
}

func getGroupBillSplits(ctx context.Context, db *sql.DB, logger *zap.Logger, billId uuid.UUID) ([]*monify.SplitPerson, error) {
	var splits []*monify.SplitPerson
	rows, err := db.QueryContext(ctx, `
		SELECT person, amount ,ui.name
		FROM group_split_bill 
		LEFT JOIN group_member gm ON group_split_bill.person = gm.group_member_id
		LEFT JOIN user_identity ui ON gm.user_id = ui.user_id
		WHERE bill_id = $1
	`, billId)
	if err != nil {
		logger.Error("", zap.Error(err))
		return splits, err
	}
	defer rows.Close()
	for rows.Next() {
		var memberId uuid.UUID
		var amount float64
		var username string
		err = rows.Scan(&memberId, &amount, &username)
		if err != nil {
			return []*monify.SplitPerson{}, err
		}
		splits = append(splits, &monify.SplitPerson{
			MemberId: memberId.String(),
			Amount:   amount,
			Username: username,
		})
	}
	return splits, nil
}

func getGroupBillPrepaid(ctx context.Context, db *sql.DB, logger *zap.Logger, billId uuid.UUID) ([]*monify.PrepaidPerson, error) {
	var prepaid []*monify.PrepaidPerson
	rows, err := db.QueryContext(ctx, `
		SELECT person, amount , ui.name
		FROM group_prepaid_bill 
		LEFT JOIN group_member gm ON group_prepaid_bill.person = gm.group_member_id
		LEFT JOIN user_identity ui ON gm.user_id = ui.user_id
		WHERE bill_id = $1
	`, billId)
	if err != nil {
		logger.Error("", zap.Error(err))
		return prepaid, err
	}
	defer rows.Close()
	for rows.Next() {
		var memberId uuid.UUID
		var amount float64
		var username string
		err = rows.Scan(&memberId, &amount, &username)
		if err != nil {
			return []*monify.PrepaidPerson{}, err
		}
		prepaid = append(prepaid, &monify.PrepaidPerson{
			MemberId: memberId.String(),
			Amount:   amount,
			Username: username,
		})
	}
	return prepaid, nil
}
