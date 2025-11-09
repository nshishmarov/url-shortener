package service

import "database/sql"

type Shortener struct {
	DB *sql.DB
}

func NewShortener(db *sql.DB) *Shortener {
	return &Shortener{DB: db}
}