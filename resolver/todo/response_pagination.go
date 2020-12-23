package todo

import (
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/model"
)

type ResponsePagination struct {
	Username string
	Data     []model.Todo `json:"data"`
}
