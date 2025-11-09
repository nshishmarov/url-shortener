package server

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"url-shortener/internal/api"
	"url-shortener/internal/config"
	"url-shortener/internal/service"
)

type Server struct {
	HttpServer *http.Server
	Port uint32
	Host string
	Log *slog.Logger
}

func NewServer(cfg *config.Config, log *slog.Logger, db *sql.DB) *Server {
	shortener := service.NewShortener(db)
	wrp := api.NewURLServerWrapper(shortener)

	mux := http.NewServeMux()
	mux.HandleFunc(cfg.Server.BaseUrl + "/url", wrp.GetURL)

	http := &http.Server{
		Addr:	fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler: mux,
	}
	
	return &Server{
		HttpServer: http,
		Port: cfg.Server.Port,
		Host: cfg.Server.Host,
		Log: log,
	}
}

func (s *Server) Start() {
	s.Log.Info("starting server on %s:%d", s.Host, s.Port)
	err := s.HttpServer.ListenAndServe()
	if err != nil {
		s.Log.Error("error starting server", err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	s.Log.Info("stopping server")
	err := s.HttpServer.Close()
	if err != nil {
		s.Log.Error("error stopping server", err)
	}
}