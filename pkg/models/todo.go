package models

import "time"

// Todo is the data model for todo tasks
type Todo struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Note      string     `json:"note"`
	DueDate   time.Time  `json:"due_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
