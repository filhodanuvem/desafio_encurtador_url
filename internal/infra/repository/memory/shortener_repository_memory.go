package memory

import "github.com/julianojj/desafio_encurtador_url/internal/domain"

type ShortenerRepositoryMemory struct {
	Shorteners []*domain.URLShortener
}

func NewShortenerRepositoryMemory() *ShortenerRepositoryMemory {
	return &ShortenerRepositoryMemory{
		Shorteners: make([]*domain.URLShortener, 0),
	}
}

func (s *ShortenerRepositoryMemory) Save(shortener *domain.URLShortener) error {
	s.Shorteners = append(s.Shorteners, shortener)
	return nil
}

func (s *ShortenerRepositoryMemory) FindByShortURL(shortURL string) (*domain.URLShortener, error) {
	var shortener *domain.URLShortener
	for _, short := range s.Shorteners {
		if short.ShortURL == shortURL {
			shortener = short
		}
	}
	return shortener, nil
}
