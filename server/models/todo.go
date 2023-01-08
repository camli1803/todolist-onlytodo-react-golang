package models

import (
	"todolist/database"

	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Content string `gorm:"column:content;not null"`
	Done    bool   `gorm:"column:done"`
}

type TodoRequest struct {
	Content string `json:"content" validate:"required,todo_content"`
	Done    bool   `json:"done"`
}

type TodoResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateTodos(todo Todo) error {
	result := database.DB.Create(&todo)
	return result.Error
}

func GetAllTodos() ([]Todo, error) {
	var todos []Todo
	result := database.DB.Find(&todos)
	return todos, result.Error
}

func GetAllTodosByStatus(done bool) (todos []Todo, err error) {
	result := database.DB.Where("done = ?", done).Find(&todos)
	return todos, result.Error
}

func GetATodoByTodoID(todoID uint64) (todo Todo, err error) {
	result := database.DB.First(&todo, todoID)
	return todo, result.Error
}

func UpdateATodoByTodoID(todoID uint64, todo Todo) error {
	result := database.DB.Model(&todo).Where("id = ?", todoID).Updates(&todo)
	return result.Error
}

func DeleteATodoByTodoID(todoID uint64) error {
	var todo Todo
	result := database.DB.Where("id = ?", todoID).Delete(&todo)
	return result.Error
}
