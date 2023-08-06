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

	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home page",
			"payload": articles,
		},
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
