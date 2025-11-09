package storage

import (
	"database/sql"
	"log/slog"
	"url-shortener/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	DB *sql.DB
	log *slog.Logger
}

func NewStorage(config *config.Config, log *slog.Logger) (*Storage, error) {
	db, err := sql.Open("pgx", config.DBConfig.DSN)
	if err != nil {
		return nil, err
	}

	return &Storage{DB: db, log: log}, nil
}

func (s *Storage) Close() error {
	return s.DB.Close()
}