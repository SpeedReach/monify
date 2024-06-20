//go:build !docker

package test

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"monify/lib/utils"
	"monify/services/api/infra"
	"os"
	"path/filepath"
	"testing"
)

func SetupTestDB() {

}
func SetupTestResource(t *testing.T, secrets map[string]string) infra.Resources {
	dbFile, err := filepath.Abs("../../../build/test.db")
	if err != nil {
		return infra.Resources{}
	}
	_ = os.Remove(dbFile)
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		panic(err.Error() + " " + dbFile)
	}
	migrationPath, err := filepath.Abs("../../../migrations")
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:"+migrationPath,
		"sqlite3", driver)
	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			panic(err)
		}
	}
	if err1, err2 := m.Close(); err1 != nil || err2 != nil {
		panic(err.Error() + err2.Error())
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	if _, err = db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		panic(err)
	}

	utils.LoadSecrets("dev")
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
