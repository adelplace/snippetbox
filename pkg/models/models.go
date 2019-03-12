package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Snippet struct {
	ID      string
	Title   string
	Content string
}
