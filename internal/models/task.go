package models

type Task struct {
	TaskID      int    `json:"task_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}