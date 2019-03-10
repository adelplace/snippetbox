package persistence

import (
	"context"
	"snippetbox/pkg/models"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// SnippetModel mongo
type SnippetModel struct {
	Collection *mongo.Collection
}

// Insert new Snippet
func (m *SnippetModel) Insert(id *primitive.ObjectID, title, content string) (*mongo.InsertOneResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := m.Collection.InsertOne(ctx, bson.M{"_id": id, "title": title, "content": content})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get a snippet
func (m *SnippetModel) Get(id string) (*models.Snippet, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var result = models.Snippet{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	err = m.Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Latest inserted Snippet
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
