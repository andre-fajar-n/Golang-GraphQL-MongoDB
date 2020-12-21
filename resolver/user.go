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

func Register(params graphql.ResolveParams) (interface{}, error) {
	username := params.Args["username"].(string)

	// check if username is already exist
	userData := repo.GetUserByUsername(context.Background(), username)
	if userData.Username != "" {
		return nil, fmt.Errorf("username has been registered")
	}

	password := params.Args["password"].(string)
	password, err := GeneratePassword(password)
	if err != nil {
		return nil, err
	}

	data := model.User{
		ID:        primitive.NewObjectID(),
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}

	if err := repo.Register(context.Background(), &data); err != nil {
		return nil, err
	}

	return data, nil
}
