package usecases

import (
	"testing"

	"github.com/julianojj/desafio_encurtador_url/internal/infra/repository/memory"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestShouldMakeANewShortenerURL(t *testing.T) {
	shortenerRepository := memory.NewShortenerRepositoryMemory()
	makeShortenerURL := NewMakeShortenerURL(shortenerRepository)
	input := MakeShortenerInputURL{
		Url: "https://github.com/julianojj",
	}
	output, err := makeShortenerURL.Execute(input)
	assert.NoError(t, err)
	shortener, err := shortenerRepository.FindByShortURL(output.ShortURL)
	assert.NoError(t, err)
	assert.Equal(t, output.ShortURL, shortener.ShortURL)
}
