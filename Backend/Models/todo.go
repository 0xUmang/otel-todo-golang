package Models

import "github.com/google/uuid"

type Todo struct {
	ID	string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`

}

func NewTodo(title string,description string,status string) *Todo {
	return &Todo{
		ID:	uuid.New().String(),
		Title: title,
		Description: description,
		Status: status,
	}
}
