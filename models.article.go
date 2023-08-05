package main

import (
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
