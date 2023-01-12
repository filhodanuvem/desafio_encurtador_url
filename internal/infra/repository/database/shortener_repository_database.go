package database

import (
	"database/sql"

	"github.com/julianojj/desafio_encurtador_url/internal/domain"
)

type ShortenerRepositoryDatabase struct {
	DB *sql.DB
}

func NewShortenerRepositoryDatabase(
	db *sql.DB,
) *ShortenerRepositoryDatabase {
	return &ShortenerRepositoryDatabase{
		DB: db,
	}
}

func (s *ShortenerRepositoryDatabase) Save(shortener *domain.URLShortener) error {
	_, err := s.DB.Exec("INSERT INTO shortener(id, longurl, shorturl, expiresin) VALUES($1, $2, $3, $4)", shortener.Id, shortener.LongURL, shortener.ShortURL, shortener.ExpiresIn)
	if err != nil {
		return err
	}
	return nil
}

func (s *ShortenerRepositoryDatabase) FindByShortURL(shortURL string) (*domain.URLShortener, error) {
	var shortener domain.URLShortener
	err := s.DB.QueryRow("SELECT id, longurl, shorturl, expiresin FROM shortener WHERE shorturl = $1", shortURL).Scan(&shortener.Id, &shortener.LongURL, &shortener.ShortURL, &shortener.ExpiresIn)
	if err != nil {
		return nil, err
	}
	return &shortener, nil
}
