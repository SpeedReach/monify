package utils

import (
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mattn/go-sqlite3"
)

func IsDuplicateKeyError(err error) bool {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		return pgErr.Code == "23505"
	}
	if sqlite, ok := err.(*sqlite3.Error); ok {
		return sqlite.ExtendedCode == sqlite3.ErrConstraintUnique
	}
	return false
}
