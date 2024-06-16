package group

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"monify/lib"
	monify "monify/protobuf/gen/go"
)

type Service struct {
	monify.UnimplementedGroupServiceServer
}

func CheckPermission(ctx context.Context, groupId uuid.UUID, userId uuid.UUID) (bool, error) {
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*) FROM group_member WHERE group_id = $1 AND user_id = $2
	`, groupId, userId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetMemberId(ctx context.Context, groupId uuid.UUID, userId uuid.UUID) (uuid.UUID, error) {
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	var memberId uuid.UUID
	rows, err := db.QueryContext(ctx, `
		SELECT group_member_id FROM group_member WHERE group_id = $1 AND user_id = $2
	`, groupId, userId)
	if err != nil {
		return uuid.Nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return uuid.Nil, nil
	}
	err = rows.Scan(&memberId)
	if err != nil {
		return uuid.Nil, err
	}
	return memberId, nil
}
