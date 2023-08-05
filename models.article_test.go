package main

import (
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	tArticleList := getAllArticles()

	if len(tArticleList) != len(articleList) {
		t.Errorf(
			"Articles count incorrect: got %d, want %d",
			len(tArticleList),
			len(articleList),
		)
	}

	for i, article := range tArticleList {
		if !(article.ID == articleList[i].ID &&
			article.Title == articleList[i].Title &&
			article.Content == articleList[i].Content) {
			t.Errorf(
				"Articles are different: got %v, want %v",
				article,
				articleList[i],
			)
		}
	}
}
