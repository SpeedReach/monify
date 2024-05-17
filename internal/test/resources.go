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
)

func SetupTestResource() infra.Resources {
	_ = os.Remove("../../build/test.db")
	db, err := sql.Open("sqlite3", "../../build/test.db")
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"sqlite3", driver)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()

	if err != nil {
		log.Fatal(err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return infra.Resources{
		DBConn: db,
		Logger: logger,
	}
}
