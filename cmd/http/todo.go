package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/binkkatal/go-todo/pkg/models"
	"github.com/binkkatal/go-todo/pkg/store"
)

// CreateTodoHandler is the handler function to create TODO
func CreateTodoHandler(todoCreator store.TodoCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todoParams models.Todo
		err := json.NewDecoder(r.Body).Decode(&todoParams)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		err = todoCreator.Create(&todoParams)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		NewJSONWriter(w).Write(todoParams, http.StatusCreated)
	}
}

// UpdateTodoHandler is the handler function to update TODO
func UpdateTodoHandler(todoUpdater store.TodoUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// skip error check as it is handled in middleware
		todo, _ := TodoFromContext(r.Context())
		var updateTodoParams models.UpdateTodoParams
		err := json.NewDecoder(r.Body).Decode(&updateTodoParams)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		if updateTodoParams.Title != nil {
			todo.Title = *updateTodoParams.Title
		}
		if updateTodoParams.Note != nil {
			todo.Note = *updateTodoParams.Note
		}
		if updateTodoParams.DueDate != nil {
			todo.DueDate = *updateTodoParams.DueDate
		}
		err = todoUpdater.Update(todo)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		NewJSONWriter(w).Write(todo, http.StatusCreated)
	}
}

// GetTodosHandler returns a list of todos
func GetTodosHandler(todoQueryer store.TodoQueryer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := map[store.TodoQueryParam]interface{}{}
		// TODO: implement params
		todos, err := todoQueryer.Query(params)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		NewJSONWriter(w).Write(todos, http.StatusOK)
	}
}

// GetTodoHandler returns a todo by their ID
func GetTodoHandler(todoGetter store.TodoGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo, _ := TodoFromContext(r.Context())
		NewJSONWriter(w).Write(todo, http.StatusOK)
	}
}

func DeleteTodoHandler(todoDeleter store.TodoDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		err := todoDeleter.Delete(id)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		NewJSONWriter(w).Write("item Deleted", http.StatusOK)
	}
}
