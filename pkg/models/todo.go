package models

import "time"

// Todo is the data model for todo tasks
type Todo struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Note    string    `json:"note"`
	DueDate time.Time `json:"due_date"`
}
