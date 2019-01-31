package persistence

import (
	"snippetbox/pkg/models"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// SnippetModel mongo
type SnippetModel struct {
	Client *mongo.Client
}

// Insert new Snippet
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get a snippet
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest inserted Snippet
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
