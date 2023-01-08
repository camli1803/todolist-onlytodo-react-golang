package main

import (
	"net/http"
	"todolist/database"
	"todolist/migrations"
	todolist_validator "todolist/validator"
	"todolist/views"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	DB, err := database.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}
	err = migrations.MigrateDB(DB)
	if err != nil {
		return
	}
	e := echo.New()

	var customValidator todolist_validator.CustomValidator
	validate := customValidator.New()
	e.Validator = &todolist_validator.CustomValidator{Validator: validate}

	e.Use(middleware.CORS())
	// api
	e.GET("/", homepage)

	todoApi := e.Group("/todos")
	{
		// todo api
		todoApi.GET("", views.GetAllTodos)
		todoApi.POST("", views.CreateTodos)
		todoApi.GET("/show_by_done", views.GetAllTodosByStatus)
		todoApi.GET("/:id", views.GetATodoByTodoID)
		todoApi.PATCH("/:id", views.UpdateATodoByTodoID)
		todoApi.DELETE("/:id", views.DeleteATodoByTodoID)
	}

	e.Logger.Fatal(e.Start(":1710"))
}

func homepage(c echo.Context) error {
	return c.JSON(http.StatusOK, "Todolist")
}
