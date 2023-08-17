// Package for interactions with MongoDB instance.
package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection instance for all DB operations.
var connectionInstance *mongo.Client
var blogDB *mongo.Database

// Check if already connected to MongoDB instance.
func isConnected(ctx context.Context) bool {
	if connectionInstance == nil {
		return false
	}
	err := connectionInstance.Ping(ctx, nil)
	if err != nil {
		return true
	}
	return true
}

// Function for connecting to MongoDB instance
// using specified MongoDB URI.
func Connect(ctx context.Context, uri string) error {
	if isConnected(ctx) {
		log.Print("Connection is already established")
		return nil
	}
	connectionInstance, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	err = connectionInstance.Ping(ctx, nil)
	if err != nil {
		return err
	}

	blogDB = connectionInstance.Database("BlogDB")
	return nil
}

// Function for closing connection with MongoDB.
func Disconnect(ctx context.Context) error {
	if isConnected(ctx) {
		log.Print("Connection doesn't exist.")
		return nil
	}

	err := connectionInstance.Disconnect(ctx)
	connectionInstance = nil
	return err
}
