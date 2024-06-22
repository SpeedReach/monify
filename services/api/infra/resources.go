package infra

import (
	"database/sql"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Resources struct {
	DBConn       *sql.DB
	Logger       *zap.Logger
	KafkaWriters KafkaWriters
	FileService  FileService
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

	grpc.NewClient("")
	return Resources{
		DBConn:       dbConn,
		Logger:       logger,
		KafkaWriters: kafkaWriters,
	}
}
