package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseLogin struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username  string
	CreatedAt time.Time
	Token     string
}
