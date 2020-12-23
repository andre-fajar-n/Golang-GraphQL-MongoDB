package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/model"
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/repo"
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/resolver/todo"
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

func ViewOneTodoByID(params graphql.ResolveParams) (interface{}, error) {
	// check token
	token := params.Context.Value("token").(string)
	verifToken, err := VerifyToken(token)
	if err != nil {
		return nil, err
	}

	id := params.Args["id"].(string)
	username := fmt.Sprintf("%v", verifToken["username"])
	data, err := repo.ViewOneTodoByID(context.Background(), id)
	if err != nil {
		return nil, err
	}

	// check jwt username is match or not with data
	if username != data.Username {
		return nil, fmt.Errorf("Unauthorized")
	}

	return data, nil
}

func ViewListTodoMe(params graphql.ResolveParams) (interface{}, error) {
	// check token
	token := params.Context.Value("token").(string)
	verifToken, err := VerifyToken(token)
	if err != nil {
		return nil, err
	}
	username := fmt.Sprintf("%v", verifToken["username"])

	page, ok := params.Args["page"].(int)
	if !ok {
		page = 1
	}
	perPage, ok := params.Args["per_page"].(int)
	if !ok {
		perPage = 10
	}

	form := repo.PaginationRequest{
		Limit:  perPage,
		Offset: (page - 1) * perPage,
	}
	data := repo.ViewListTodoMe(context.Background(), username, form)

	return todo.ResponsePagination{
		Username: username,
		Data:     data,
	}, nil
}
