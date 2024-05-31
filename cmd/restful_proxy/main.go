package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
	_ = gw.RegisterGroupsBillServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)

	handler := CorsHandler(mux)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", handler)
}

func CorsHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		h.ServeHTTP(w, r)
	})
}
