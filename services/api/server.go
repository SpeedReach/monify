package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"monify/lib"
	authLib "monify/lib/auth"
	monify "monify/protobuf/gen/go"
	"monify/services/api/controllers/auth"
	"monify/services/api/controllers/group"
	"monify/services/api/controllers/group_bill"
	"monify/services/api/infra"
	"net"
)

type Server struct {
	server    *grpc.Server
	Config    ServerConfig
	Resources infra.Resources
}

func NewServer(config ServerConfig, resources infra.Resources) Server {
	g := grpc.NewServer(
		setupInterceptor(resources, config),
		grpc.MaxRecvMsgSize(16*1024*1024), // 16 MB, adjust as needed
		grpc.MaxSendMsgSize(16*1024*1024), // 16 MB, adjust as needed
	)
	SetupServices(g, config)
	return Server{
		server:    g,
		Config:    config,
		Resources: resources,
	}
}

func SetupServices(g *grpc.Server, config ServerConfig) {
	monify.RegisterAuthServiceServer(g, auth.Service{
		Secret: config.JwtSecret,
	})
	monify.RegisterGroupServiceServer(g, group.Service{})
	monify.RegisterGroupsBillServiceServer(g, group_bill.Service{})
}

func setupInterceptor(resources infra.Resources, config ServerConfig) grpc.ServerOption {
	authFn := authLib.AuthMiddleware{JwtSecret: config.JwtSecret}
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		userId, err := authFn.GrpcExtractUserId(ctx, req, info)
		if err != nil {
			return nil, err
		}
		if userId != uuid.Nil {
			ctx = context.WithValue(ctx, lib.UserIdContextKey{}, userId)
		}
		requestId := uuid.New()
		ctx = context.WithValue(ctx, lib.KafkaWriterContextKey{}, resources.KafkaWriters)
		ctx = context.WithValue(ctx, lib.DatabaseContextKey{}, resources.DBConn)
		ctx = context.WithValue(ctx, lib.LoggerContextKey{}, resources.Logger.With(zap.String("request_id", requestId.String())))
		m, err := handler(ctx, req)
		return m, err
	}
	return grpc.UnaryInterceptor(interceptor)
}

func (s Server) Start(lis net.Listener) error {
	if err := s.server.Serve(lis); err != nil {
		return errors.New(fmt.Sprintf("failed to serve: %v", err))
	}
	return nil
}