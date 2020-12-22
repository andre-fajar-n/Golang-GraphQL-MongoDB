package repo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/infrastructure"
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/model"
)

func AddTodo(ctx context.Context, data *model.Todo) error {
	_, err := infrastructure.Mongodb.Collection("todo").InsertOne(ctx, data)
	return err
}

func ViewOneTodoByID(ctx context.Context, id, username string) model.Todo {
	var todo model.Todo

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	data := infrastructure.Mongodb.Collection("todo").FindOne(ctx, bson.M{
		"_id":      objectId,
		"username": username,
	})
	data.Decode(&todo)

	return todo
}
