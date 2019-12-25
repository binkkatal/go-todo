package api

import (
	"github.com/binkkatal/go-todo/pkg/models"
	"github.com/binkkatal/go-todo/pkg/store"
	"github.com/binkkatal/go-todo/pkg/utils"
)

// TodoService performs operations on todo items.
type TodoService struct {
	TodoStore store.TodoStore
	Logger    utils.Logger
}

// Get returns an Todo.
func (svc *TodoService) Get(id string) (*models.Todo, error) {
	todo, err := svc.TodoStore.Get(id)
	if err != nil {
		if err == utils.ErrorNotFound {
			return nil, ErrorNotFound(`todo item not found.`)
		}

		go svc.Logger.Error(err)
		return nil, ErrorInternal(`Internal error.`)
	}

	return todo, nil
}

// Query returns Todos.
func (svc *TodoService) Query(params map[store.TodoQueryParam]interface{}) ([]models.Todo, error) {

	todos, err := svc.TodoStore.Query(params)
	if err != nil {
		go svc.Logger.Error(err)
		return nil, ErrorInternal(`Internal error.`)
	}

	return todos, nil
}

// Create Creates a Todo.
func (svc *TodoService) Create(todo *models.Todo) error {
	tx := svc.TodoStore.BeginTx()

	err := tx.Create(todo)
	if err != nil {
		go svc.Logger.Error(err)
		return ErrorInternal(`Internal error.`)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return ErrorInternal(`Internal error.`)
	}
	return nil
}

// Update method updates the todo item
func (svc *TodoService) Update(todo *models.Todo) error {
	tx := svc.TodoStore.BeginTx()

	err := tx.Update(todo)
	if err != nil {
		if err == utils.ErrorInternal {
			return ErrorNotFound(`Unable to Update todo item.`)
		}

		go svc.Logger.Error(err)
		return ErrorInternal(`Internal error.`)
	}

	err = tx.Commit()
	if err != nil {
		return ErrorInternal(`Internal error.`)
	}

	return nil
}

// Delete Deletes a todo item.
func (svc *TodoService) Delete(id string) error {
	tx := svc.TodoStore.BeginTx()

	err := tx.Delete(id)

	if err != nil {
		if err == utils.ErrorInternal {
			return ErrorNotFound(`Unable to Delete Todo as the record is not found.`)
		}

		go svc.Logger.Error(err)
		return ErrorInternal(`Internal error.`)
	}

	err = tx.Commit()
	if err != nil {
		return ErrorInternal(`Internal error.`)
	}
	return nil
}
