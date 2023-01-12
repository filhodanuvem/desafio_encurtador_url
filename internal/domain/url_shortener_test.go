package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNotShouldCreateANewShortenerIfExpiredURL(t *testing.T) {
	expiresIn := time.Now().AddDate(0, 0, -1)
	shortener, err := NewURLShortener(uuid.NewString(), expiresIn)
	assert.EqualError(t, err, "expired url")
	assert.Nil(t, shortener)
}

func TestNotShouldShortenIfURLIsRequired(t *testing.T) {
	expiresIn := time.Now().AddDate(0, 0, 1)
	shortener, _ := NewURLShortener(uuid.NewString(), expiresIn)
	shortenedURL, err := shortener.Shorten("")
	assert.EqualError(t, err, "url is required")
	assert.Empty(t, shortenedURL)
}

func TestShouldCreateANewShortener(t *testing.T) {
	expiresIn := time.Now().AddDate(0, 0, 1)
	shortener, err := NewURLShortener(uuid.NewString(), expiresIn)
	shortener.Shorten("https://github.com/julianojj")
	assert.Equal(t, "https://github.com/julianojj", shortener.LongURL)
	assert.NotEmpty(t, shortener.ShortURL)
	assert.NoError(t, err)
}
