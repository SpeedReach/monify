package media

import (
	"context"
	"fmt"
	"log"
	"monify/lib"
	"monify/lib/auth"
	"net"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer(infra Infra) Server {
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
	mux.Handle("/image", resourceMiddleware(authMiddleware.HttpMiddleware(http.HandlerFunc(uploadImage))))
	return Server{
		mux: mux,
	}
}

func (s Server) Start() {
	port := "8080"
	//start listening
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
	err = http.Serve(lis, s.mux)
	if err != nil {
		return
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
