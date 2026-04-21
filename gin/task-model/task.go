package taskmodel

type Task struct{
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}