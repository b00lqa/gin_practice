package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Creating router.
	router := gin.Default()
	// Loading templates.
	router.LoadHTMLGlob("templates/*")
	// Defining router handler.
	router.GET("/", showIndexPage)
	router.Run()
}
