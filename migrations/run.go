package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"log"
	"monify/internal/utils"
)

func main() {

	_ = godotenv.Load()
	secrets, err := utils.LoadSecrets(utils.LoadEnv())
	if err != nil {
		panic(secrets)
	}

	db, err := sql.Open("pgx", secrets["POSTGRES_URI"])
	if err != nil {
		log.Fatal(err)
	}
	driver, err := cockroachdb.WithInstance(db, &cockroachdb.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://.",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err = m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}
