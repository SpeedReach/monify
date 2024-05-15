package internal

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"monify/internal/infra"
	"monify/internal/middlewares"
	"net"
)

type Server struct {
	server *grpc.Server
	Config infra.Config
}

func NewServer(config infra.Config) Server {
	s := Server{
		server: grpc.NewServer(setupInterceptor(config)),
		Config: config,
	}
	return s
}

func setupInterceptor(config infra.Config) grpc.ServerOption {
	authFn := middlewares.AuthMiddleware{JwtSecret: config.JwtSecret}
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		err := authFn.PreHandler(ctx, req, info)
		if err != nil {
			return nil, err
		}

		m, err := handler(ctx, req)

		return m, err
	}
	return grpc.UnaryInterceptor(interceptor)
}

func (s Server) Start(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
