package repo

import (
	"context"

	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/infrastructure"
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/model"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(ctx context.Context, data *model.User) error {
	_, err := infrastructure.Mongodb.Collection("user").InsertOne(ctx, data)
	return err
}

func GetUserByUsername(ctx context.Context, username string) model.User {
	var user model.User
	data := infrastructure.Mongodb.Collection("user").FindOne(ctx, bson.M{"username": username})
	data.Decode(&user)
	return user
}
