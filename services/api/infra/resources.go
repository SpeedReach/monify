package infra

import (
	"database/sql"
	"go.uber.org/zap"
)

type Resources struct {
	DBConn       *sql.DB
	Logger       *zap.Logger
	KafkaWriters KafkaWriters
}

func SetupResources(config Config) Resources {
	dbConn := SetupConnection(config.PostgresURI)
	kafkaWriters, err := NewKafkaWriter(config)
	if err != nil {
		panic(err)
	}
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return Resources{
		DBConn:       dbConn,
		Logger:       logger,
		KafkaWriters: kafkaWriters,
	}
}
