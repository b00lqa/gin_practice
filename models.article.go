package main

import (
	"errors"

	"github.com/google/uuid"
)

// Article struct containing all info about articles.
type article struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

// TODO add db support.
var articleList = []article{
	{ID: uuid.New(), Title: "Article 1", Content: "Blank content for testing only."},
	{ID: uuid.New(), Title: "Article 2", Content: "Blank content for testing only."},
}

// Function for retreiving all articles.
func getAllArticles() []article {
	return articleList
}

// Function for getting article by it's id.
func getArticleByID(id uuid.UUID) (*article, error) {
	for _, article := range articleList {
		if id == article.ID {
			return &article, nil
		}
	}

	return nil, errors.New("article not found")
}
