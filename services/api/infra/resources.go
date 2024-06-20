package infra

import (
	"database/sql"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	monify "monify/protobuf/gen/go"
)

type Resources struct {
	DBConn       *sql.DB
	Logger       *zap.Logger
	KafkaWriters KafkaWriters
	ImageService *monify.MediaServiceClient
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
