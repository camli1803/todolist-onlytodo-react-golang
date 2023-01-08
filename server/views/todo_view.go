package views

import (
	"net/http"
	"strconv"
	"strings"
	"todolist/controllers"
	"todolist/models"

	"github.com/labstack/echo/v4"
)

func CreateTodos(c echo.Context) error {
	var todoCreateRequest models.TodoRequest
	err := c.Bind(&todoCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(todoCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = controllers.CreateTodos(todoCreateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "todo has been created")
}

func GetAllTodos(c echo.Context) error {
	todos, err := controllers.GetAllTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}

func GetAllTodosByStatus(c echo.Context) error {
	done := c.QueryParam("done")
	var donebool bool
	if strings.EqualFold(done, "true") {
		donebool = true
	} else if strings.EqualFold(done, "false") {
		donebool = false
	} else {
		return c.JSON(http.StatusBadRequest, "done must be true or false")
	}

	todosRes, err := controllers.GetAllTodosByStatus(donebool)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todosRes)
}

func GetATodoByTodoID(c echo.Context) error {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todoRes, err := controllers.GetATodoByTodoID(todoID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todoRes)

}

func UpdateATodoByTodoID(c echo.Context) error {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var todoUpdateRequest models.TodoRequest
	err = c.Bind(&todoUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(todoUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = controllers.UpdateATodoByTodoID(todoID, todoUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "todo has been updated")
}

func DeleteATodoByTodoID(c echo.Context) error {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = controllers.DeleteATodoByTodoID(todoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "todo has been deleted")
}