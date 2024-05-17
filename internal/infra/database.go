package infra

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func SetupConnection(uri string) *sql.DB {
	db, err := sql.Open("pgx", uri)
	if err != nil {
		panic(err)
	}
	return db
}
