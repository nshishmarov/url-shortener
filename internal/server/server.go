package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"url-shortener/internal/api"
	"url-shortener/internal/config"
	"url-shortener/internal/service"
	"url-shortener/internal/storage"
)

type Server struct {
	HttpServer *http.Server
	Port uint32
	Host string
	Log *slog.Logger
}

func NewServer(cfg *config.Config, log *slog.Logger, storage *storage.Storage) *Server {
	shortener := service.NewShortener(storage)
	wrp := api.NewURLServerWrapper(shortener)

	mux := http.NewServeMux()
	mux.HandleFunc(cfg.Server.BaseUrl + "/url", wrp.HandleUrl)

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