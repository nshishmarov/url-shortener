package api

import (
	"net/http"
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

func (u *URLServerWrapper) GetURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Hello, world!"}`))
}