package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/model"
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/repo"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTodo(params graphql.ResolveParams) (interface{}, error) {
	// check token
	token := params.Context.Value("token").(string)
	verifToken, err := VerifyToken(token)
	if err != nil {
		return nil, err
	}

	task := params.Args["task"].(string)
	deadline := params.Args["deadline"].(string)
	layout := "2006-01-02T15:04:05Z"
	deadlineTime, err := time.Parse(layout, deadline)
	if err != nil {
		str := fmt.Sprintf("Invalid format date, e.g.: %s", layout)
		return nil, fmt.Errorf(str)
	}

	data := model.Todo{
		ID:        primitive.NewObjectID(),
		Task:      task,
		Deadline:  deadlineTime,
		Username:  fmt.Sprintf("%v", verifToken["username"]),
		IsDone:    false,
		CreatedAt: time.Now(),
	}

	if err = repo.AddTodo(context.Background(), &data); err != nil {
		return nil, err
	}

	return data, nil
}
