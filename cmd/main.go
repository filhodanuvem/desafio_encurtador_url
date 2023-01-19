package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julianojj/desafio_encurtador_url/internal/infra/api/controllers"
	"github.com/julianojj/desafio_encurtador_url/internal/infra/api/routes"
	"github.com/julianojj/desafio_encurtador_url/internal/infra/repository/database"
	"github.com/julianojj/desafio_encurtador_url/internal/usecases"
	_ "github.com/lib/pq"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// database
	connectionString := os.Getenv("BASE_URL")
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
		r,
		makeShortenerController,
		getOriginalURLController,
	).Init()

	// log.Println("Starting server on port: ", port)
	// server.ListenAndServe()
	r.Run(":3000")
}
