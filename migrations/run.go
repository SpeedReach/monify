package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"monify/internal/utils"
)

func main() {

	_ = godotenv.Load()
	secrets, err := utils.LoadSecrets(utils.LoadEnv())
	if err != nil {
		panic(secrets)
	}
	m, err := migrate.New(
		"file://.",
		secrets["POSTGRES_URI"])
	if err != nil {

		panic(err)
	}
	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			panic(err)
		} else {
			println("No Change")
		}
	}
}
