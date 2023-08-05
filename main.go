package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creating router
	router := gin.Default()
	// Loading templates
	router.LoadHTMLGlob("templates/*")
	// Defining router handler
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})
	router.Run()
}
