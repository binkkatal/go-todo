package models

import "time"

// UpdateTodoParams is the data model for todo tasks
type UpdateTodoParams struct {
	Title   *string    `json:"title"`
	Note    *string    `json:"note"`
	DueDate *time.Time `json:"due_date"`
}
