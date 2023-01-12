package usecases

import (
	"database/sql"
	"log"
	"testing"

	"github.com/julianojj/desafio_encurtador_url/internal/infra/repository/database"
	"github.com/stretchr/testify/assert"
)

func TestSHouldCreateANewShortenerAndGetOriginalURLBasedOnShortUrl(t *testing.T) {
	connectionString := "postgres://juliano:123456@localhost:5432/app?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	shortenerRepository := database.NewShortenerRepositoryDatabase(db)
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
