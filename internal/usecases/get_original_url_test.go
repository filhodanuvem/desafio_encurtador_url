package usecases

import (
	"testing"

	"github.com/julianojj/desafio_encurtador_url/internal/infra/repository/memory"
	"github.com/stretchr/testify/assert"
)

func TestSHouldCreateANewShortenerAndGetOriginalURLBasedOnShortUrl(t *testing.T) {
	shortenerRepository := memory.NewShortenerRepositoryMemory()
	makeShortenerURL := NewMakeShortenerURL(shortenerRepository)
	inputMakeShortenerURL := MakeShortenerInputURL{
		Url: "https://github.com/julianojj",
	}
	output, err := makeShortenerURL.Execute(inputMakeShortenerURL)
	assert.NoError(t, err)
	getOriginalURL := NewGetOriginalURL(shortenerRepository)
	shortener, err := getOriginalURL.Execute(output.ShortURL)
	assert.NoError(t, err)
	assert.Equal(t, inputMakeShortenerURL.Url, shortener.Url)
}
