package group_bill

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
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
	row := db.QueryRowContext(ctx, `
		SELECT group_id FROM group_member WHERE group_member_id = $1
	`, history.operator)

	var groupId uuid.UUID
	if err := row.Scan(&groupId); err != nil {
		return err
	}

	_, err := db.ExecContext(ctx, `
		INSERT INTO group_bill_history( history_id, type, bill_id, title, operator, group_id) VALUES ($1, $2, $3, $4, $5, $6)
	`, uuid.New(), history.ty, history.billId, history.title, history.operator, groupId)
	return err
}

func getBillGroupId(ctx context.Context, db *sql.DB, billId uuid.UUID) (uuid.UUID, error) {
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
func getBillGroupIdAndCreator(ctx context.Context, db *sql.DB, billId uuid.UUID) (uuid.UUID, uuid.UUID, error) {
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
