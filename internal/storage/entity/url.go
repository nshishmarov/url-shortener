package entity

import "github.com/google/uuid"

type URL struct {
	ID  uuid.UUID
	URL string
	ShortURL string
}