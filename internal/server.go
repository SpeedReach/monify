package internal

import (
	"monify/internal/handlers"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewProduction() Server {
	s := Server{
		mux: http.NewServeMux(),
	}
	s.setupHandlers()
	return s
}

func (s Server) setupHandlers() {
	s.mux.HandleFunc("/", handlers.ExampleHandler)
}

func (s Server) Start(port string) {
	err := http.ListenAndServe(":"+port, s.mux)
	if err != nil {
		return
	}

}
