package api

import (
	"encoding/json"
	"net/http"
	"url-shortener/internal/api/dto"
	"url-shortener/internal/api/mapper"
	"url-shortener/internal/service"
)

type URLServerWrapperInterface interface {
	GetURL()
}

type URLServerWrapper struct {
	Shortener *service.Shortener
}

func NewURLServerWrapper(shortener *service.Shortener) *URLServerWrapper {
	return &URLServerWrapper{Shortener: shortener}
}

func (u *URLServerWrapper) HandleUrl(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.GetURL(w, r)
	case http.MethodPost:
		u.CreateURL(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (u *URLServerWrapper) GetURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shortURL := r.URL.Query().Get("short_url")
	longURL, err :=u.Shortener.GetLongURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	longUrlDto := mapper.MapUrlEntityToUrlDto(longURL)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(longUrlDto)
}

func (u *URLServerWrapper) CreateURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var longUrlDto dto.URL
	err := json.NewDecoder(r.Body).Decode(&longUrlDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortUrlDto, err := u.Shortener.CreateShortURL(&longUrlDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shortUrlDto)
}