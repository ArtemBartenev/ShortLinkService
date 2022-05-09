package repository

import (
	"shortlink/internals/domain/entity"
)

type LinkRepository interface {
	Add(link *entity.Link) error
	FindShortURLByOriginalURL(url string) (string, error)
	FindOriginalURLByShortURL(url string) (string, error)
}
