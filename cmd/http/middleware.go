package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/binkkatal/go-todo/pkg/api"
	"github.com/binkkatal/go-todo/pkg/models"
)

// ContextKey is a key for a context value.
type ContextKey string

const (
	contextKeyTodo         = ContextKey("todo")
)



// ContextWithTodo returns a context with todo.
func ContextWithTodo(ctx context.Context, todo *models.Todo) context.Context {
	return context.WithValue(ctx, contextKeyTodo, todo)
}

// TodoFromContext returns a Todo from the given context.
func TodoFromContext(ctx context.Context) (*models.Todo, bool) {
	if v := ctx.Value(contextKeyTodo); v != nil {
		if todo, ok := v.(*models.Todo); ok {
			return todo, ok
		}
	}
	return nil, false
}
// TodoMiddlewareHandler is  a middleware for todo service api to get todo item
func TodoMiddlewareHandler(svc *api.TodoService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			todo, err := svc.Get(id)
			if err != nil {
				NewErrorWriter(w).Write(err)
				return
			}
			ctx := ContextWithTodo(r.Context(), todo)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}