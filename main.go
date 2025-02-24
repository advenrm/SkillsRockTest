package main

import (
	"fmt"
	postgres "skillsrocktest/database"
	"skillsrocktest/handler"
	"skillsrocktest/repository"
	"skillsrocktest/routes"
	"skillsrocktest/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type ApiResponse struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	taskRepo := repository.NewPostgreSQLRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	port := 3000
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		response := ApiResponse{Data: "data", Error: ""}
		return c.JSON(response)
	})

	app.Get("/template", func(c *fiber.Ctx) error {
		return c.Render("example", fiber.Map{
			"Name": "Roman Morozov",
		})
	})

	routes.SetRoutes(app, taskHandler)

	app.Listen((fmt.Sprintf(":%d", port)))
}
