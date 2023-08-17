package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Article struct containing all info about articles.
type Article struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title   string             `json:"title" bson:"title,omitempty"`
	Content string             `json:"content" bson:"content,omitempty"`
}

// Function for retreiving all articles.
func GetAllArticles(ctx context.Context) ([]Article, error) {
	coll := blogDB.Collection("articles")
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return []Article{}, err
	}

	var articleList []Article
	err = cursor.All(ctx, &articleList)
	if err != nil {
		return []Article{}, err
	}

	return articleList, nil
}

// Function for getting article by it's id.
func GetArticleByID(ctx context.Context, id_str string) (*Article, error) {
	id, err := primitive.ObjectIDFromHex(id_str)
	if err != nil {
		return nil, err
	}

	coll := blogDB.Collection("articles")
	article_mongo := coll.FindOne(ctx, bson.M{"_id": id})
	if article_mongo == nil {
		return nil, errors.New("article not found")
	}

	var article Article
	article_mongo.Decode(&article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}
