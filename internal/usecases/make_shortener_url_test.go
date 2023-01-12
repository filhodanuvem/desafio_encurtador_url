package usecases

import (
	"database/sql"
	"log"
	"testing"

	"github.com/julianojj/desafio_encurtador_url/internal/infra/repository/database"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestShouldMakeANewShortenerURL(t *testing.T) {
	connectionString := "postgres://juliano:123456@localhost:5432/app?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	shortenerRepository := database.NewShortenerRepositoryDatabase(db)
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
