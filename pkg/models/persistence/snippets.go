package persistence

import (
	"context"
	"time"

	"github.com/adelplace/snippetbox/pkg/models"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// SnippetModel mongo
type SnippetModel struct {
	Collection *mongo.Collection
}

// Insert new Snippet
func (m *SnippetModel) Insert(id *primitive.ObjectID, title, content string) (*mongo.InsertOneResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := m.Collection.InsertOne(ctx, bson.M{"_id": id, "title": title, "content": content, "created": time.Now()})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get a snippet
func (m *SnippetModel) Get(id primitive.ObjectID) (*models.Snippet, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var result = models.Snippet{}
	filter := bson.M{"_id": id}
	err := m.Collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, models.ErrNoRecord
	}
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Latest inserted Snippet
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{}
	options := options.FindOptions{Sort: bson.M{"id": -1}}
	result, err := m.Collection.Find(ctx, filter, &options)
	if err != nil {
		return nil, err
	}

	for result.Next(ctx) {

	}

	return nil, nil
}
