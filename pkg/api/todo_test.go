package api_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/binkkatal/go-todo/pkg/api"
	"github.com/binkkatal/go-todo/pkg/models"
	"github.com/binkkatal/go-todo/pkg/postgres"
	"github.com/binkkatal/go-todo/pkg/store"
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	todoColumns = []string{"id", "title", "note", "due_date", "created_at", "updated_at", "deleted_at"}
	currentTime = time.Now()
	todo1       = models.Todo{
		ID:        ksuid.New().String(),
		Note:      "this is a test note",
		Title:     "This is a test title",
		CreatedAt: currentTime,
	}
	todo2 = models.Todo{
		ID:        ksuid.New().String(),
		Note:      "this is a test note",
		Title:     "This is a test title",
		CreatedAt: currentTime,
	}
)

func TestGetTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectQuery(
		regexp.QuoteMeta(postgres.FindTodo)).WithArgs(todo1.ID).WillReturnRows(
		sqlmock.NewRows(todoColumns).AddRow(todo1.ID, todo1.Title, todo1.Note, todo1.DueDate, todo1.CreatedAt, todo1.UpdatedAt, todo1.DeletedAt))
	defer db.Close()

	svc := api.TodoService{
		TodoStore: &postgres.TodoDB{
			DB: &postgres.DB{db},
		},
		Logger: logrus.StandardLogger(),
	}
	item, err := svc.Get(todo1.ID)
	assert.NoError(t, err, "Should get record by ID")
	assert.Equal(t, todo1.ID, item.ID, "Should get expected ID for record")
	assert.Equal(t, todo1.Note, item.Note, "Should get expected ID for record")
	assert.Equal(t, todo1.Title, item.Title, "Should get expected Title for record")
	assert.Equal(t, todo1.CreatedAt, item.CreatedAt, "Should get expected CreatedAt for record")
	assert.Equal(t, todo1.DueDate, item.DueDate, "Should get expected DueDate for record")

}

func TestQueryTodo(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectQuery(
		regexp.QuoteMeta(fmt.Sprintf(postgres.QueryTodos, "todos.deleted_at IS NULL"))).WillReturnRows(
		sqlmock.NewRows(todoColumns).AddRow(
			todo1.ID, todo1.Title, todo1.Note, todo1.DueDate, todo1.CreatedAt, todo1.UpdatedAt, todo1.DeletedAt,
		).AddRow(
			todo2.ID, todo2.Title, todo2.Note, todo2.DueDate, todo2.CreatedAt, todo2.UpdatedAt, todo2.DeletedAt,
		))

	defer db.Close()

	svc := api.TodoService{
		TodoStore: &postgres.TodoDB{
			DB: &postgres.DB{db},
		},
		Logger: logrus.StandardLogger(),
	}
	todos, err := svc.Query(map[store.TodoQueryParam]interface{}{})

	assert.NoError(t, err, "Should query records without error")
	assert.Len(t, todos, 2, "Should get 2 results")
	todo := todos[1]
	assert.Equal(t, todo2.ID, todo.ID, "Should get expected ID for record")
	assert.Equal(t, todo2.Note, todo.Note, "Should get expected ID for record")
	assert.Equal(t, todo2.Title, todo.Title, "Should get expected Title for record")
	assert.Equal(t, todo2.CreatedAt, todo.CreatedAt, "Should get expected CreatedAt for record")
	assert.Equal(t, todo2.DueDate, todo.DueDate, "Should get expected DueDate for record")
}

func TestCreateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(postgres.CreateTodo)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	defer db.Close()

	svc := api.TodoService{
		TodoStore: &postgres.TodoDB{
			DB: &postgres.DB{db},
		},
		Logger: logrus.StandardLogger(),
	}
	err = svc.Create(&todo1)
	assert.NoError(t, err, "should not error on create")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "Should meet expectations")
}

func TestUpdateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(postgres.UpdateTodo)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	defer db.Close()

	svc := api.TodoService{
		TodoStore: &postgres.TodoDB{
			DB: &postgres.DB{db},
		},
		Logger: logrus.StandardLogger(),
	}
	err = svc.Update(&todo1)
	assert.NoError(t, err, "should not error on update")
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "Should meet expectations")
}

func TestDeleteTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(postgres.DeleteTodo)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	defer db.Close()

	svc := api.TodoService{
		TodoStore: &postgres.TodoDB{
			DB: &postgres.DB{db},
		},
		Logger: logrus.StandardLogger(),
	}
	err = svc.Delete(todo1.ID)
	assert.NoError(t, err, "should not error on delete")
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "Should meet expectations")
}
