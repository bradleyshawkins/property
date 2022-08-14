package rest

import (
	"context"
	"github.com/bradleyshawkins/property/moving"
	"github.com/bradleyshawkins/property/postgres"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type Server struct {
	mux      *chi.Mux
	moveIn   moving.MoveIn
	database *postgres.Database
}

func NewServer(database *postgres.Database) *Server {
	mux := chi.NewRouter()

	s := &Server{
		mux:      mux,
		database: database,
	}

	s.RegisterRoutes()

	return s
}

func (s *Server) RegisterRoutes() {
	s.mux.Post("/resident/{residentID}/movein", s.MoveIn)
	s.mux.Post("/resident/{residentID}/claim", nil)
}

func (s *Server) Start(port string) func(ctx context.Context) error {
	srv := http.Server{
		Addr:    ":" + port,
		Handler: s.mux,
	}

	go func() {
		log.Println("Starting http server ...")

		err := http.ListenAndServe(":"+port, s.mux)
		if err != nil {
			log.Println("Error shutting down server. Error:", err)
		}
	}()

	return func(ctx context.Context) error {
		log.Println("Shutting down http server ...")
		return srv.Shutdown(ctx)
	}
}
