package main

import (
	"github.com/binkkatal/go-todo/pkg/api"
	"github.com/go-chi/chi"
)

type API struct {
	TodoService *api.TodoService
}

// APIMux returns an API multiplexer.
func APIMux(api *API) *chi.Mux {
	mux := chi.NewMux()

	mux.Route("/todo", func(r chi.Router) {
		r.Post("/", CreateTodoHandler(api.TodoService))
		r.Get("/", GetTodosHandler(api.TodoService))
		r.Route("/{id}", func(r chi.Router) {
			r.With(TodoMiddlewareHandler(api.TodoService)).Get("/", GetTodoHandler(api.TodoService))
			r.With(TodoMiddlewareHandler(api.TodoService)).Delete("/", DeleteTodoHandler(api.TodoService))
			r.With(TodoMiddlewareHandler(api.TodoService)).Put("/", UpdateTodoHandler(api.TodoService))
		})
	})
	return mux
}
