package repo

import (
	"context"

	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/infrastructure"
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/model"
)

func AddTodo(ctx context.Context, data *model.Todo) error {
	_, err := infrastructure.Mongodb.Collection("todo").InsertOne(ctx, data)
	return err
}
