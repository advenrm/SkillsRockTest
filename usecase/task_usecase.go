package usecase

import (
	"skillsrocktest/models"
	"skillsrocktest/repository"
)

type TaskUsecase interface {
	AddTask(title, description, status string) error
	GetTasks() ([]models.Task, error)
	UpdateTask(id uint, title, description, status string) error
	DeleteTask(id uint) error
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{
		taskRepo: taskRepo,
	}
}

func (uc *taskUsecase) AddTask(title, description, status string) error {
	task := models.Task{
		Title:       title,
		Description: description,
		Status:      status,
	}

	return uc.taskRepo.Add(task)
}

func (uc *taskUsecase) GetTasks() ([]models.Task, error) {
	return uc.taskRepo.GetList()
}

func (uc *taskUsecase) UpdateTask(id uint, title, description, status string) error {
	task := models.Task{
		Id:          id,
		Title:       title,
		Description: description,
		Status:      status,
	}
	return uc.taskRepo.UpdateTaskByID(task)

}

func (uc *taskUsecase) DeleteTask(id uint) error {
	return uc.taskRepo.DeleteTaskByID(id)

}
