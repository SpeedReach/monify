package infra

import (
	"database/sql"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	monify "monify/protobuf/gen/go"
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

	conn, err := grpc.NewClient("media_service:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return Resources{
		DBConn:       dbConn,
		Logger:       logger,
		KafkaWriters: kafkaWriters,
		FileService: FileService{
			config: config,
			client: monify.NewMediaServiceClient(conn),
		},
	}
}
