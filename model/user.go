package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserBaseModel struct {
	Username string
	Password string
}

type UserWithID struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string
	Password string
}
