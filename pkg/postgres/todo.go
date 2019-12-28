package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/binkkatal/go-todo/pkg/models"
	"github.com/binkkatal/go-todo/pkg/store"
	"github.com/binkkatal/go-todo/pkg/utils"
	"github.com/segmentio/ksuid"
)

const (
	FindTodo = `
		SELECT
			todos.id,
			todos.title,
			todos.note,
			todos.due_date,
			todos.created_at,
			todos.updated_at,
			todos.deleted_at
		from
			todos
		WHERE
			todos.id = $1 AND todos.deleted_at IS NULL`
	QueryTodos = `
		SELECT
			todos.id,
			todos.title,
			todos.note,
			todos.due_date,
			todos.created_at,
			todos.updated_at,
			todos.deleted_at
		from
			todos
		WHERE
			%s
		ORDER BY
			todos.due_date`

	CreateTodo = `
		INSERT INTO todos (
			id,
			title,
			note,
			due_date
		)
		VALUES ($1, $2, $3, $4)
	`

	UpdateTodo = `
		UPDATE todos
		SET
			title=$2,
			due_date=$3,
			note=$4,
			updated_at=$5
		WHERE id = $1
	`

	DeleteTodo = `
		UPDATE todos
		SET deleted_at=$2
		WHERE id = $1
	`
)

// TodoDB is a database for Todos.
type TodoDB struct {
	*DB
}

// TodoTx is an Todo transaction.
type TodoTx struct {
	*Tx
}

// Get returns an Todo.
func (db *TodoDB) Get(id string) (*models.Todo, error) {
	var todo models.Todo

	err := db.QueryRow(FindTodo, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Note,
		&todo.DueDate,
		&todo.CreatedAt,
		&todo.UpdatedAt, &todo.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrorNotFound
		}
		return nil, err
	}
	return &todo, nil
}

// Query returns a list of Todos.
func (db *TodoDB) Query(params map[store.TodoQueryParam]interface{}) ([]models.Todo, error) {
	query := QueryTodos
	wheres := []string{`todos.deleted_at IS NULL`}

	if params != nil {
		// TODO: add parameters for query
	}

	query = fmt.Sprintf(query, strings.Join(wheres, " AND "))
	// TODO: add parameters for query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos := []models.Todo{}
	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Note,
			&todo.DueDate,
			&todo.CreatedAt,
			&todo.UpdatedAt, &todo.DeletedAt,
		)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

// BeginTx begins a transaction.
func (db *TodoDB) BeginTx() store.TodoTx {
	return &TodoTx{Tx: db.Begin()}
}

// Create creates an Todo.
func (tx *TodoTx) Create(todo *models.Todo) error {

	todo.ID = ksuid.New().String()
	todo.CreatedAt = time.Now()

	query := CreateTodo
	parameters := make([]interface{}, 4)
	parameters[0] = todo.ID
	parameters[1] = todo.Title
	parameters[2] = todo.Note
	parameters[3] = todo.DueDate
	_, err := tx.Exec(query, parameters...)
	if err != nil {
		return err
	}
	return nil
}

// Update updates an Todo.
func (tx *TodoTx) Update(todo *models.Todo) error {
	t := time.Now()
	todo.UpdatedAt = &t

	query := UpdateTodo

	parameters := make([]interface{}, 5)
	parameters[0] = todo.ID
	parameters[1] = todo.Title
	parameters[2] = todo.DueDate
	parameters[3] = todo.Note
	parameters[4] = todo.UpdatedAt
	_, err := tx.Exec(query, parameters...)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes an Todo.
func (tx *TodoTx) Delete(id string) error {
	query := `
		UPDATE todos
		SET deleted_at=$2
		WHERE id = $1
	`

	_, err := tx.Exec(query, id, time.Now())
	if err != nil {
		return err
	}

	return nil
}
