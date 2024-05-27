package group

import (
	"context"
	"database/sql"
	monify "monify/protobuf/gen/go"

	"github.com/google/uuid"
)

type Service struct {
	monify.UnimplementedGroupServiceServer
}

func createGroupMember(ctx context.Context, db *sql.DB, groupId uuid.UUID, userId uuid.UUID) (uuid.UUID, error) {
	memberId := uuid.New()
	_, err := db.Exec(`
		INSERT INTO group_member (group_id, user_id) VALUES ($1,$2)
	`, groupId, userId)
	if err != nil {

		return uuid.Nil, err
	}
	return memberId, err
}
