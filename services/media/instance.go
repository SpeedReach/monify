package media

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"monify/lib"
	"monify/lib/auth"
	monify "monify/protobuf/gen/go"
	"net"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	mux     *http.ServeMux
	gServer *grpc.Server
}

func NewServer(infra Infra) Server {
	return Server{
		mux:     NewHttpServer(infra),
		gServer: NewGrpcServer(infra),
	}
}

func NewHttpServer(infra Infra) *http.ServeMux {
	mux := http.NewServeMux()
	authMiddleware := auth.AuthMiddleware{JwtSecret: infra.config.JwtSecret}
	resourceMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, lib.ImageStorageContextKey{}, infra.objStorage)
			ctx = context.WithValue(ctx, lib.DatabaseContextKey{}, infra.db)
			ctx = context.WithValue(ctx, lib.ConfigContextKey{}, infra.config)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	mux.Handle("/image", allowCORS(resourceMiddleware(authMiddleware.HttpMiddleware(http.HandlerFunc(uploadImage)))))
	return mux
}

func NewGrpcServer(infra Infra) *grpc.Server {
	s := grpc.NewServer(setupInterceptor(infra))
	monify.RegisterMediaServiceServer(s, Service{})
	return s
}

func (s Server) Start() {
	port1 := "8080"
	port2 := "8081"
	//start listening
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port2))
		if err != nil {
			panic(err)
		}
		log.Fatal(s.gServer.Serve(lis))
	}()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port1))
	if err != nil {
		panic(err)
	}
	log.Fatal(http.Serve(lis, s.mux))
}

func setupInterceptor(resources Infra) grpc.ServerOption {
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		requestId := uuid.New()
		logger := resources.logger.With(zap.String("request_id", requestId.String()))
		ctx = context.WithValue(ctx, lib.DatabaseContextKey{}, resources.db)
		ctx = context.WithValue(ctx, lib.LoggerContextKey{}, logger)
		ctx = context.WithValue(ctx, lib.ImageStorageContextKey{}, resources.objStorage)
		start := time.Now()
		m, err := handler(ctx, req)
		logger.Log(zap.InfoLevel, "request completed", zap.Duration("duration", time.Since(start)), zap.String("method", info.FullMethod))
		return m, err
	}
	return grpc.UnaryInterceptor(interceptor)
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		preflightHandler(w, r)
		if r.Method == http.MethodOptions {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	grpclog.Infof("Preflight request for %s", r.URL.Path)
}
