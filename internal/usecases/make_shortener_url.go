package usecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/julianojj/desafio_encurtador_url/internal/domain"
	"github.com/julianojj/desafio_encurtador_url/internal/repository"
)

type MakeShortenerURL struct {
	ShortenerRepository repository.ShortenerRepository
}

type MakeShortenerInputURL struct {
	Url string `json:"url"`
}

type MakeShortenerOutputURL struct {
	ShortURL string `json:"short"`
}

func NewMakeShortenerURL(shortenerRepository repository.ShortenerRepository) *MakeShortenerURL {
	return &MakeShortenerURL{
		ShortenerRepository: shortenerRepository,
	}
}

func (m *MakeShortenerURL) Execute(input MakeShortenerInputURL) (MakeShortenerOutputURL, error) {
	expiresIn := time.Now().AddDate(0, 0, 1)
	shortener, err := domain.NewURLShortener(uuid.NewString(), expiresIn)
	if err != nil {
		return MakeShortenerOutputURL{}, err
	}
	shortenedURL, err := shortener.Shorten(input.Url)
	if err != nil {
		return MakeShortenerOutputURL{}, err
	}
	err = m.ShortenerRepository.Save(shortener)
	if err != nil {
		return MakeShortenerOutputURL{}, err
	}
	return MakeShortenerOutputURL{
		ShortURL: shortenedURL,
	}, nil
}
