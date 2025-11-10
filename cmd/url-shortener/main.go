package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"
	"url-shortener/internal/config"
	"url-shortener/internal/logger"
	"url-shortener/internal/server"
	"url-shortener/internal/storage"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Mode)

	log.Info("starting url-shortener")

	storage, err := storage.NewStorage(cfg, log)
	if err != nil {
		log.Error("failed to create storage", "err", err)
		return
	}
	defer func() {
		err := storage.Close()
		if err != nil {
			log.Error("failed to close storage", "err", err)
		}
	}()

	ctx , cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	server := server.NewServer(cfg, log, storage)

	go func() {
		<-ctx.Done()
		
		ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
		defer cancel()

		server.Stop(ctx)
	}()

	server.Start()
}

