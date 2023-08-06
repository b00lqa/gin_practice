package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
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
