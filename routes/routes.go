package routes

import (
	"skillsrocktest/handler"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App, taskHandler *handler.TaskHandler) {
	tasksAPI := app.Group("/tasks")

	tasksAPI.Static("/", "./static")
	tasksAPI.Static("/edit", "./static")

	tasksAPI.Post("/", taskHandler.AddTask)
	tasksAPI.Get("/", taskHandler.GetTasks)
	tasksAPI.Get("/edit/:id", taskHandler.EditTask)
	tasksAPI.Put("/:id", taskHandler.UpdateTask)
	tasksAPI.Delete("/:id", taskHandler.DeleteTask)
}
