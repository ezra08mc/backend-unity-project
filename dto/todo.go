package dto

import "time"

type TodoRequest struct {
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

type TodoResponse struct {
	Success     bool       `json:"success"`
	Message     string     `json:"message"`
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	IsDone      bool       `json:"is_done"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
