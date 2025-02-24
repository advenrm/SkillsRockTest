package handler

import (
	"skillsrocktest/models"
	"skillsrocktest/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		taskUsecase: taskUsecase,
	}
}

func (h *TaskHandler) AddTask(c *fiber.Ctx) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	status := c.FormValue("status")

	err := h.taskUsecase.AddTask(title, description, status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add a new task")
	}

	return c.Redirect("/tasks")
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	tasks, err := h.taskUsecase.GetTasks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve tasks")
	}

	return c.Render("index", fiber.Map{
		"Tasks": tasks,
	})
}

func (h *TaskHandler) EditTask(c *fiber.Ctx) error {
	id := c.Params("id")

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
	}

	task := &models.Task{
		Id: uint(taskID),
	}

	return c.Render("edit", fiber.Map{"Task": task})
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	title := c.FormValue("title")
	description := c.FormValue("description")
	status := c.FormValue("status")

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
	}

	if err := h.taskUsecase.UpdateTask(uint(taskID), title, description, status); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update task")
	}

	return c.Redirect("/tasks")
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
	}

	if err := h.taskUsecase.DeleteTask(uint(taskID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete task")
	}

	return c.Redirect("/tasks")
}
