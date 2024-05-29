package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	gw "monify/protobuf/gen/go"
	"net/http"
)

var grpcServerEndpoint = flag.String("grpc-server-endpoint", ":2302", "gRPC server endpoint")

func main() {
	flag.Parse()
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	println("grpcServerEndpoint: ", *grpcServerEndpoint)
	err := gw.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	_ = gw.RegisterGroupServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", handler)
}
