//go:build docker

package test

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
	"monify/internal/infra"
	"testing"
)

func SetupTestResource(t *testing.T) infra.Resources {
	db, err := sql.Open("pgx", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			panic(err)
		}
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	return infra.Resources{
		DBConn: db,
		Logger: logger,
	}
}
