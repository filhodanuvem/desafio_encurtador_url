package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/julianojj/desafio_encurtador_url/internal/infra/api/controllers"
	"github.com/julianojj/desafio_encurtador_url/internal/infra/api/routes"
	"github.com/julianojj/desafio_encurtador_url/internal/infra/repository/database"
	"github.com/julianojj/desafio_encurtador_url/internal/usecases"
	_ "github.com/lib/pq"
)

func main() {
	mux := http.NewServeMux()
	port := 3000
	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", port),
		Handler: mux,
	}

	// database
	connectionString := "postgres://juliano:123456@localhost:5432/app?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// repositories
	shortenerRepository := database.NewShortenerRepositoryDatabase(db)
	getOriginalURL := usecases.NewGetOriginalURL(shortenerRepository)
	makeShortenerURL := usecases.NewMakeShortenerURL(shortenerRepository)

	// controllers
	makeShortenerController := controllers.NewShortenerController(makeShortenerURL)
	getOriginalURLController := controllers.NewGetOriginalURLController(getOriginalURL)

	// routes
	routes.NewShortenerRoute(
		mux,
		makeShortenerController,
		getOriginalURLController,
	).Init()

	log.Println("Starting server on port: ", port)
	server.ListenAndServe()
}
