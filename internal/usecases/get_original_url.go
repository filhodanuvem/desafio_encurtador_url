package usecases

import (
	"errors"

	"github.com/julianojj/desafio_encurtador_url/internal/repository"
)

type GetOriginalURL struct {
	ShortenerRepository repository.ShortenerRepository
}

type GetOriginalURLOutput struct {
	Url string `json:"url"`
}

func NewGetOriginalURL(shortenerRepository repository.ShortenerRepository) *GetOriginalURL {
	return &GetOriginalURL{
		ShortenerRepository: shortenerRepository,
	}
}

func (g *GetOriginalURL) Execute(shortUrl string) (GetOriginalURLOutput, error) {
	existingShortenerURL, err := g.ShortenerRepository.FindByShortURL(shortUrl)
	if err != nil {
		return GetOriginalURLOutput{}, err
	}
	if existingShortenerURL == nil {
		return GetOriginalURLOutput{}, errors.New("short not found")
	}
	if existingShortenerURL.IsExpired() {
		return GetOriginalURLOutput{}, errors.New("expired short")
	}
	return GetOriginalURLOutput{
		Url: existingShortenerURL.LongURL,
	}, nil
}
