package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/binkkatal/go-todo/pkg/api"
	"github.com/binkkatal/go-todo/pkg/postgres"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	// PostgresQL (DB)
	db, err := postgres.Open(DBConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	todoService := &api.TodoService{
		TodoStore: &postgres.TodoDB{
			DB: db,
		},
		Logger: logrus.StandardLogger(),
	}

	api := &API{TodoService: todoService}
	mux := chi.NewMux()
	mux.Mount("/api/v1", APIMux(api))
	ListenAndServe(MustGetenv("PORT"), mux)
}

// ListenAndServe runs the server.
func ListenAndServe(port string, handler http.Handler) {
	fmt.Println("Listening on:", MustGetenv("PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler); err != nil {
		panic(err)
	}
}

// DBConnectionString returns the database connection string.
func DBConnectionString() string {
	connectionString := fmt.Sprintf(`host=%s`, MustGetenv("DB_HOST"))

	if port := os.Getenv("DB_PORT"); port != "" {
		connectionString += fmt.Sprintf(` port=%s`, port)
	}

	if user := os.Getenv("DB_USER"); user != "" {
		connectionString += fmt.Sprintf(` user=%s`, user)
	}

	if password := os.Getenv("DB_PASSWORD"); password != "" {
		connectionString += fmt.Sprintf(` password=%s`, password)
	}

	if name := os.Getenv("DB_NAME"); name != "" {
		connectionString += fmt.Sprintf(` dbname=%s`, name)
	}

	if mode := os.Getenv("DB_SSL_MODE"); mode != "" {
		connectionString += fmt.Sprintf(` sslmode=%s`, mode)
	}

	return connectionString
}

// MustGetenv gets an environment variable or panics.
func MustGetenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		logrus.Panicf("%s missing",key)
	}
	return v
}
