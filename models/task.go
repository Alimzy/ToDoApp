package models

import "time"

// Status represents the current state of a task
type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

// Task is the core data structure for our API
type Task struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreateTaskInput holds the fields a client sends when creating a task
type CreateTaskInput struct {
	Title       string `json:"title"       binding:"required"`
	Description string `json:"description"`
}
