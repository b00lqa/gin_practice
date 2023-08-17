package main

import (
	"net/http"

	"github.com/b00lqa/gin_practice/mongodb"
	"github.com/gin-gonic/gin"
)

// Function handler for processing index.html
// with articles.
func showIndexPage(ctx *gin.Context) {
	articles, err := mongodb.GetAllArticles(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	render(
		ctx,
		gin.H{
			"title":   "Home page",
			"payload": articles,
		},
		"index.html",
	)
}

// Function handler for viewing article.
func getArticle(ctx *gin.Context) {
	article_id := ctx.Param("article_id")

	article, err := mongodb.GetArticleByID(ctx, article_id)
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

// Render one of HTML, JSON or XML based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present.
func render(ctx *gin.Context, data gin.H, templateName string) {

	switch ctx.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON.
		ctx.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML.
		ctx.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML.
		ctx.HTML(http.StatusOK, templateName, data)
	}

}
