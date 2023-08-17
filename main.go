package main

import (
	"context"
	"log"
	"text/template"

	"github.com/b00lqa/gin_practice/mongodb"
	"github.com/caarlos0/env/v9"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Structure containing all configuration variables
// passed in .env file or system environment.
type Config struct {
	MongoDB_URI string `env:"MONGODB_URI"`
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
		log.Fatalf("Couldn't parse environment variables: %v", err)
	}

	log.Println("Connecting to db...")
	// Initialize MongoDB connection.
	err = mongodb.Connect(context.Background(), cfg.MongoDB_URI)
	if err != nil {
		log.Fatalf("Couldn't connect to db: %v", err)
	}
	// Defer closing connection to MongoDB
	defer mongodb.Disconnect(context.Background())

	// Creating router.
	router := gin.Default()
	// Map router functions.
	router.SetFuncMap(template.FuncMap{
		// Function for converting ObjectIDs to Hex representation
		// so it can be used in URL.
		"ObjectIDToHex": func(id primitive.ObjectID) string {
			return id.Hex()
		},
	})
	// Loading templates.
	router.LoadHTMLGlob("templates/*")
	// Defining index router handler.
	router.GET("/", showIndexPage)
	// Defining article route handler.
	router.GET("/article/view/:article_id", getArticle)

	router.Run()
}
