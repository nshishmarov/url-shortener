package service

import (
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"url-shortener/internal/api/dto"
	"url-shortener/internal/storage"
	"url-shortener/internal/storage/entity"
)

const (
	baseShortUrl = "short.ty"
)

type Shortener struct {
	storage *storage.Storage
	log *slog.Logger
}

func NewShortener(storage *storage.Storage) *Shortener {
	return &Shortener{storage: storage, log: storage.Log}
}

func (s *Shortener) GetLongURL(shortURL string) (*entity.URL, error) {
	longURL, err := s.storage.GetLongUrl(shortURL)
	if err != nil {
		s.log.Error("failed to get long url", err)
		return nil, err
	}

	return longURL, nil
}

func (s *Shortener) CreateShortURL(longURL *dto.URL) (*entity.URL, error) {
	genShortUrl, err := s.generateShortURL(longURL)
	if err != nil {
		s.log.Error("failed to generate short url", err)
		return nil, err
	}

	shortURL, err := s.storage.CreateShortUrl(longURL, genShortUrl)
	if err != nil {
		s.log.Error("failed to create short url", err)
		return nil, err
	}

	return shortURL, nil
}

func (s *Shortener) generateShortURL(longURL *dto.URL) (*dto.URL, error) {
	hash := sha256.Sum256([]byte(longURL.URL))
	enc := hex.EncodeToString(hash[:5])
	return &dto.URL{URL: baseShortUrl + "/" + enc}, nil
}