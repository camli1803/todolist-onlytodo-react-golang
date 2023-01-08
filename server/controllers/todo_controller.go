package controllers

import (
	"todolist/models"
)

func CreateTodos(todoCreateRequest models.TodoRequest) error {
	var todo models.Todo
	todo.Content = todoCreateRequest.Content
	todo.Done = todoCreateRequest.Done

	return models.CreateTodos(todo)
}

func GetAllTodos() (todoRes []models.TodoResponse, err error) {
	todos, err := models.GetAllTodos()
	if err != nil {
		return nil, err
	}

	for _, todo := range todos {
		todoRes = append(todoRes, convertTodoEntityToTodoResponse(todo))
	}
	return todoRes, nil
}

func GetAllTodosByStatus(done bool) (todosRes []models.TodoResponse, err error) {
	todos, err := models.GetAllTodosByStatus(done)
	if err != nil {
		return nil, err
	}
	for _, todo := range todos {
		todosRes = append(todosRes, convertTodoEntityToTodoResponse(todo))
	}
	return todosRes, nil
}

func GetATodoByTodoID(todoID uint64) (todoRes models.TodoResponse, err error) {
	todo, err := models.GetATodoByTodoID(todoID)
	todoRes = convertTodoEntityToTodoResponse(todo)
	return todoRes, err
}

func UpdateATodoByTodoID(todoID uint64, todoUpdateRequest models.TodoRequest) error {
	var todo models.Todo
	todo.Content = todoUpdateRequest.Content
	todo.Done = todoUpdateRequest.Done
	return models.UpdateATodoByTodoID(todoID, todo)
}

func DeleteATodoByTodoID(todoID uint64) error {
	return models.DeleteATodoByTodoID(todoID)
}

func convertTodoEntityToTodoResponse(todo models.Todo) (todoResponse models.TodoResponse) {
	todoResponse.ID = todo.ID
	todoResponse.Content = todo.Content
	todoResponse.Done = todo.Done
	todoResponse.CreatedAt = todo.CreatedAt
	todoResponse.UpdatedAt = todo.UpdatedAt
	return todoResponse
}
