package service

import (
	"fmt"
	"shortlink/internals/domain/entity"
	"shortlink/internals/repository"
	"shortlink/pkg/encoding"
)

type Converter interface {
	ConvertOriginalURLToShortURL(link *entity.Link) (string, error)
	GetOriginalURLByShortURL(link *entity.Link) (string, error)
}

type LinkService struct {
	repository repository.LinkRepository
}

func NewLinkService(repository repository.LinkRepository) *LinkService {
	return &LinkService{repository: repository}
}

func (s *LinkService) ConvertOriginalURLToShortURL(link *entity.Link) (string, error) {
	var ShortURL string

	ShortURL, err := s.repository.FindShortURLByOriginalURL(link.OriginalURL)

	if err != nil {
		return "", err
	}

	if ShortURL != "" {
		return "", fmt.Errorf("Original URL already exists.")
	}

	link.ShortURL += encoding.GenerateStringByBase62()
	s.repository.Add(link)

	return ShortURL, nil
}

func (s *LinkService) GetOriginalURLByShortURL(link *entity.Link) (string, error) {
	OriginalURL, err := s.repository.FindOriginalURLByShortURL(link.ShortURL)

	if err != nil {
		return "", err
	}

	return OriginalURL, nil
}
