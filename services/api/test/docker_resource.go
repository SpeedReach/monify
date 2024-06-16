//go:build docker

package test

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
	"monify/services/api/infra"
	"path/filepath"
	"testing"
)

func SetupTestResource(t *testing.T, secrets map[string]string) infra.Resources {
	db, err := sql.Open("pgx", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	migrationPath, err := filepath.Abs("../../../migrations")
	m, err := migrate.NewWithDatabaseInstance(
		"file:"+migrationPath,
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

	kafWriters, err := infra.NewKafkaWriter(infra.NewConfigFromSecrets(secrets))
	if err != nil {
		panic(err)
	}
	return infra.Resources{
		DBConn:       db,
		Logger:       logger,
		KafkaWriters: kafWriters,
	}
}
