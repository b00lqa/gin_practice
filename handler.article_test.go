package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := io.ReadAll(w.Body)
		pageOK := err == nil &&
			strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

func TestArticleUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/view/:article_id", getArticle)

	articles := getAllArticles()
	for _, article := range articles {
		req, _ := http.NewRequest("GET", "/article/view/"+article.ID.String(), nil)
		testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
			statusOK := w.Code == http.StatusOK

			p, err := io.ReadAll(w.Body)
			pageOK := err == nil &&
				strings.Index(string(p), article.Title) > 0

			return statusOK && pageOK
		})
	}
}
