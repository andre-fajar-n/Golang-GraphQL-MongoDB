package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Task      string
	IsDone    bool
	Username  string
	Deadline  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
