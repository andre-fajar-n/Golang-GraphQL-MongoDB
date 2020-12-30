package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/infrastructure"
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PaginationRequest struct {
	Limit  int
	Offset int
}

func AddTodo(ctx context.Context, data *model.Todo) error {
	_, err := infrastructure.Mongodb.Collection("todo").InsertOne(ctx, data)
	return err
}

func ViewOneTodoByID(ctx context.Context, id string) (model.Todo, error) {
	var todo model.Todo

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return todo, fmt.Errorf("Invalid ID")
	}

	filter := bson.M{"_id": objectId}
	data := infrastructure.Mongodb.Collection("todo").FindOne(ctx, filter)
	data.Decode(&todo)

	return todo, nil
}

func ViewListTodoMe(ctx context.Context, username string, form PaginationRequest) []model.Todo {
	var todo model.Todo
	var todos []model.Todo

	option := options.Find().
		SetLimit(int64(form.Limit)).
		SetSkip(int64(form.Offset))
	filter := bson.M{
		"username": username,
	}
	data, err := infrastructure.Mongodb.Collection("todo").Find(ctx, filter, option)
	defer data.Close(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}

	for data.Next(ctx) {
		data.Decode(&todo)
		todos = append(todos, todo)
	}

	return todos
}
