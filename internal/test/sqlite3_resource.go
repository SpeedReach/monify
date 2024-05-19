//go:build !docker

package test

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"log"
	"monify/internal/infra"
	"os"
	"testing"
)

func SetupTestResource(t *testing.T) infra.Resources {

	_ = os.Remove("../../build/test.db")
	db, err := sql.Open("sqlite3", "../../build/test.db")
	if err != nil {
		log.Fatal(err)
	}
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"sqlite3", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err = m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if err1, err2 := m.Close(); err1 != nil || err2 != nil {
		log.Fatal(err, err2)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	db, err = sql.Open("sqlite3", "../../build/test.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	return infra.Resources{
		DBConn: db,
		Logger: logger,
	}
}
