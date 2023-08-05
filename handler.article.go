package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Function handler for processing index.html
// with articles.
func showIndexPage(ctx *gin.Context) {
	articles := getAllArticles()

	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home page",
			"payload": articles,
		},
	)

}
