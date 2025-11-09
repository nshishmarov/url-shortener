package main

import (
	"url-shortener/internal/config"
	"url-shortener/internal/logger"
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
}

