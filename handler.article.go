package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Function handler for processing index.html
// with articles.
func showIndexPage(ctx *gin.Context) {
	articles := getAllArticles()

	render(
		ctx,
		gin.H{
			"title":   "Home page",
			"payload": articles,
		},
		"index.html",
	)
}

func getArticle(ctx *gin.Context) {
	article_id, err := uuid.Parse(ctx.Param("article_id"))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}

	article, err := getArticleByID(article_id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.HTML(
		http.StatusOK,
		"article.html",
		gin.H{
			"payload": article,
		},
	)
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(ctx *gin.Context, data gin.H, templateName string) {

	switch ctx.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		ctx.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		ctx.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		ctx.HTML(http.StatusOK, templateName, data)
	}

}
