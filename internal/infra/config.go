package infra

import (
	"monify/internal/utils"
	"os"
)

type Config struct {
	Environment string
	PostgresURI string
	JwtSecret   string
}

func NewProductionConfig() Config {
	environment, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		environment = "dev"
	}

	secrets, err := utils.LoadSecrets(environment)
	if err != nil {
		panic(err)
	}

	return Config{
		Environment: environment,
		PostgresURI: secrets["POSTGRES_URI"],
		JwtSecret:   secrets["JWT_SECRET"],
	}

}
