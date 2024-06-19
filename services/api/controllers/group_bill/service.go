package group_bill

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"monify/lib"
	"monify/lib/group_bill"
	monify "monify/protobuf/gen/go"
	"monify/services/api/infra"
)

type Service struct {
	monify.UnimplementedGroupsBillServiceServer
}

func processGroupBillModifyEvent(ctx context.Context, db *sql.Tx, history group_bill.GroupBillModification) error {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	if _, err := db.ExecContext(ctx, `
		INSERT INTO group_bill_history( history_id, type, bill_id, title, operator, group_id) VALUES ($1, $2, $3, $4, $5, $6)
	`, uuid.New(), history.Ty, history.BillId, history.Title, history.OperatorMemberId, history.GroupId); err != nil {
		logger.Error("failed to insert bill history", zap.Error(err))
		return err
	}
	kfWriter := ctx.Value(lib.KafkaWriterContextKey{}).(infra.KafkaWriters)
	serialized, err := json.Marshal(history)
	if err != nil {
		logger.Error("", zap.Error(err))
		return err
	}
	if err = kfWriter.GroupBill.WriteMessages(ctx,
		kafka.Message{Value: serialized},
	); err != nil {
		logger.Error("", zap.Error(err))
		return err
	}
	return nil
}

// context requires middlewares.DatabaseContextKey:*sql.DB
func getBillGroupId(ctx context.Context, billId uuid.UUID) (uuid.UUID, error) {
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
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
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
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
