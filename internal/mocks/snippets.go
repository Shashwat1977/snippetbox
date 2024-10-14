package mocks

import (
	"time"

	"github.com.Shashwat1977.snippetbox/internal/models"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "Test Snippet",
	Content: "Test Content",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 1, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {

	return []*models.Snippet{mockSnippet}, nil

}
