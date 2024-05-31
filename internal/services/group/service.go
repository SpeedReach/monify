package group

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	monify "monify/protobuf/gen/go"
)

type Service struct {
	monify.UnimplementedGroupServiceServer
}

func CheckPermission(ctx context.Context, db *sql.DB, groupId uuid.UUID, userId uuid.UUID) (bool, error) {
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*) FROM group_member WHERE group_id = $1 AND user_id = $2
	`, groupId, userId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
