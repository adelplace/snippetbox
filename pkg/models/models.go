package models

import (
	"errors"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Snippet struct {
	ID      primitive.ObjectID `bson:"_id"`
	Title   string
	Content string
	Created time.Time
}
