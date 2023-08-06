package main

import (
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Structure containing all configuration variables
// passed in .env file or system environment.
type Config struct {
	MongoDB_URL string `env:"MONGODB_URL"`
}

func main() {
	// Load env variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Couldn't load environment variables: %v", err)
	}

	cfg := Config{}
	// Parsing environment variables to structure.
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Coutdn't parse environment variables: %v", err)
	}

	// Creating router.
	router := gin.Default()
	// Loading templates.
	router.LoadHTMLGlob("templates/*")
	// Defining index router handler.
	router.GET("/", showIndexPage)
	// Defining article route handler.
	router.GET("/article/view/:article_id", getArticle)

	router.Run()
}
