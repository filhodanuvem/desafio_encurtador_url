package repository

import "github.com/julianojj/desafio_encurtador_url/internal/domain"

type ShortenerRepository interface {
	Save(shortener *domain.URLShortener) error
	FindByShortURL(shorturl string) (*domain.URLShortener, error)
}
