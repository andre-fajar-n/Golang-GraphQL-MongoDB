package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Task      string
	IsDone    bool `bson:"is_done" json:"is_done"`
	Username  string
	Deadline  time.Time
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
