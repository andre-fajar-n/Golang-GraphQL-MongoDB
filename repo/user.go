package repo

import (
	"context"

	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/infrastructure"
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/model"
	"go.mongodb.org/mongo-driver/bson"
)

// var collection = infrastructure.Mongodb.Collection("user")

func Register(ctx context.Context, data *model.UserBaseModel) error {
	_, err := infrastructure.Mongodb.Collection("user").InsertOne(ctx, data)
	return err
}

func GetUserByUsername(ctx context.Context, username string) model.UserWithID {
	var user model.UserWithID
	data := infrastructure.Mongodb.Collection("user").FindOne(ctx, bson.M{"username": username})
	data.Decode(&user)
	return user
}
