package resolver

import (
	"context"
	"fmt"

	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/model"
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/repo"
	"github.com/graphql-go/graphql"
)

func Register(params graphql.ResolveParams) (interface{}, error) {
	username := params.Args["username"].(string)
	data := model.UserBaseModel{
		Username: username,
		Password: params.Args["password"].(string),
	}

	userData := repo.GetUserByUsername(context.Background(), username)
	if userData.Username != "" {
		return nil, fmt.Errorf("username has been registered")
	}

	if err := repo.Register(context.Background(), &data); err != nil {
		return nil, err
	}

	return data, nil
}
