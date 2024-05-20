package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"monify/internal/infra"
	"monify/internal/middlewares"
	"monify/internal/services/auth"
	"monify/internal/services/group"
	monify "monify/protobuf"
	"net"
)

type Server struct {
	server    *grpc.Server
	Config    ServerConfig
	Resources infra.Resources
}

func NewServer(config ServerConfig, resources infra.Resources) Server {
	g := grpc.NewServer(setupInterceptor(resources, config))
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
}

func setupInterceptor(resources infra.Resources, config ServerConfig) grpc.ServerOption {
	authFn := middlewares.AuthMiddleware{JwtSecret: config.JwtSecret}
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		userId, err := authFn.ExtractUserId(ctx, req, info)
		if err != nil {
			return nil, err
		}
		if userId != uuid.Nil {
			ctx = context.WithValue(ctx, middlewares.UserIdContextKey{}, userId)
		}
		ctx = context.WithValue(ctx, middlewares.StorageContextKey{}, resources.DBConn)
		ctx = context.WithValue(ctx, middlewares.LoggerContextKey{}, resources.Logger)
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
