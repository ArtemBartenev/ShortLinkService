package entity

import "github.com/google/uuid"

type Link struct {
	ID          uuid.UUID
	ShortURL    string
	OriginalURL string
}

func NewLink() *Link {
	return &Link{
		ID: uuid.New(),
	}
}
