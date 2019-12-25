package store

import (
	"github.com/binkkatal/go-todo/pkg/models"
)

// TodoQueryParam is an Todo query parameter.
type TodoQueryParam string

// TodoQueryParamLimit indicates the maximum number of todo's to return.
const TodoQueryParamLimit = TodoQueryParam("limit")

// TodoGetter is the interface that wraps a Todo get request.
type TodoGetter interface {
	Get(id string) (*models.Todo, error)
}

// TodoQueryer is the interface that wraps a Todo query request.
type TodoQueryer interface {
	Query(params map[TodoQueryParam]interface{}) ([]models.Todo, error)
}

// TodoCreator is the interface that wraps a Todo creation request.
type TodoCreator interface {
	Create(todo *models.Todo) error
}

// TodoUpdater is the interface that wraps a Todo update request.
type TodoUpdater interface {
	Update(todo *models.Todo) error
}

// TodoDeleter is the interface that wraps a Todo delete request.
type TodoDeleter interface {
	Delete(id string) error
}

// TodoTxBeginner is the interface that wraps a Todo transaction starter.
type TodoTxBeginner interface {
	BeginTx() TodoTx
}

// TodoStore defines the operations of a Todo store.
type TodoStore interface {
	TodoGetter
	TodoQueryer
	TodoTxBeginner
}

// TodoTx defines the operations that may be performed on a Todo update transaction.
type TodoTx interface {
	TodoCreator
	TodoUpdater
	TodoDeleter
	Commit() error
	Rollback() error
}
