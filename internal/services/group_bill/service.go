package group_bill

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"monify/internal/middlewares"
	monify "monify/protobuf/gen/go"
)

type Service struct {
	monify.UnimplementedGroupsBillServiceServer
}

type HistoryType = int

const (
	Create HistoryType = 0
	Delete HistoryType = 1
	Modify HistoryType = 2
)

type billHistoryInsertion struct {
	ty       HistoryType
	operator uuid.UUID
	billId   uuid.UUID
	title    string
}

func insertBillHistory(ctx context.Context, db *sql.Tx, history billHistoryInsertion) error {
	logger := ctx.Value(middlewares.LoggerContextKey{}).(*zap.Logger)
	row := db.QueryRowContext(ctx, `
		SELECT group_id FROM group_member WHERE group_member_id = $1
	`, history.operator)

	var groupId uuid.UUID
	if err := row.Scan(&groupId); err != nil {
		logger.Error("failed to get group_id", zap.Error(err))
		return err
	}

	if _, err := db.ExecContext(ctx, `
		INSERT INTO group_bill_history( history_id, type, bill_id, title, operator, group_id) VALUES ($1, $2, $3, $4, $5, $6)
	`, uuid.New(), history.ty, history.billId, history.title, history.operator, groupId); err != nil {
		logger.Error("failed to insert bill history", zap.Error(err))
		return err
	}

	return nil
}

// context requires middlewares.DatabaseContextKey:*sql.DB
func getBillGroupId(ctx context.Context, billId uuid.UUID) (uuid.UUID, error) {
	db := ctx.Value(middlewares.DatabaseContextKey{}).(*sql.DB)
	row := db.QueryRowContext(ctx, `
		SELECT group_id FROM group_bill WHERE bill_id = $1
	`, billId)

	var groupId uuid.UUID
	if err := row.Scan(&groupId); err != nil {
		if err == sql.ErrNoRows {
			return uuid.Nil, nil
		}
		return uuid.Nil, err
	}
	return groupId, nil
}

// Returns group_id then created_by
// context requires middlewares.DatabaseContextKey:*sql.DB
func getBillGroupIdAndCreator(ctx context.Context, billId uuid.UUID) (uuid.UUID, uuid.UUID, error) {
	db := ctx.Value(middlewares.DatabaseContextKey{}).(*sql.DB)
	row := db.QueryRowContext(ctx, `
		SELECT group_id, created_by FROM group_bill WHERE bill_id = $1
	`, billId)

	var groupId uuid.UUID
	var createdBy uuid.UUID
	if err := row.Scan(&groupId, &createdBy); err != nil {
		if err == sql.ErrNoRows {
			return uuid.Nil, uuid.Nil, nil
		}
		return uuid.Nil, uuid.Nil, err
	}
	return groupId, createdBy, nil
}
