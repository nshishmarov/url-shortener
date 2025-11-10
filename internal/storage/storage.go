package storage

import (
	"database/sql"
	"log/slog"
	"url-shortener/internal/api/dto"
	"url-shortener/internal/config"
	"url-shortener/internal/storage/entity"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	DB *sql.DB
	Log *slog.Logger
}

func NewStorage(config *config.Config, log *slog.Logger) (*Storage, error) {
	db, err := sql.Open("pgx", config.DBConfig.DSN)
	if err != nil {
		return nil, err
	}

	return &Storage{DB: db, Log: log}, nil
}

func (s *Storage) Close() error {
	return s.DB.Close()
}

func (s *Storage) GetLongUrl(shortUrl string) (*entity.URL, error) {
	err := s.DB.Ping()
	if err != nil {
		s.Log.Error("failed to ping database", err)
	}
	
	row := s.DB.QueryRow(
		"SELECT * FROM urls WHERE short_url = $1",
		shortUrl,
	)
	if row.Err() != nil {
		s.Log.Error("failed to query database", row.Err().Error())
		return nil, row.Err()
	}

	var longURL entity.URL
	row.Scan(&longURL.ID, &longURL.URL, &longURL.ShortURL)

	return &longURL, nil
}

func (s *Storage) CreateShortUrl(longUrl *dto.URL, shortUrlStr *dto.URL) (*entity.URL, error) {
	err := s.DB.Ping()
	if err != nil {
		s.Log.Error("failed to ping database", err)
	}

	row := s.DB.QueryRow(
		"INSERT INTO urls (url, short_url) VALUES ($1, $2) RETURNING id, url, short_url",
		longUrl.URL,
		shortUrlStr.URL,
	)
	if row.Err() != nil {
		s.Log.Error("failed to query database", row.Err().Error())
		return nil, row.Err()
	}

	var shortURL entity.URL
	row.Scan(&shortURL.ID, &shortURL.URL, &shortURL.ShortURL)

	return &shortURL, nil
}