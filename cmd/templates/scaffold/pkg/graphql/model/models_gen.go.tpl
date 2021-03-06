// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// A task that is waiting to be finished.
type Todo struct {
	// A unique string for identifying the task.
	ID string `json:"id"`
	// The details of the task.
	Description string `json:"description"`
	// The username who created the task.
	Username string `json:"username"`
	// The creation time of the task.
	CreatedAt time.Time `json:"createdAt"`
}
