package infra

import (
	"database/sql"
	"go.uber.org/zap"
)

type Resources struct {
	DBConn *sql.DB
	Logger *zap.Logger
}

func SetupResources(config Config) Resources {
	dbConn := SetupConnection(config.PostgresURI)
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return Resources{
		DBConn: dbConn,
		Logger: logger,
	}
}
