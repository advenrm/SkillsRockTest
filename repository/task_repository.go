package repository

import (
	"database/sql"
	"fmt"
	"skillsrocktest/models"
)

type TaskRepository interface {
	Add(models.Task) error
	GetList() ([]models.Task, error)
	UpdateTaskByID(task models.Task) error
	DeleteTaskByID(id uint) error
}

type PostgreSQLRepository struct {
	db *sql.DB
}

func NewPostgreSQLRepository(db *sql.DB) *PostgreSQLRepository {
	return &PostgreSQLRepository{
		db: db,
	}
}

func (r *PostgreSQLRepository) Add(task models.Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)", task.Title, task.Description, task.Status)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgreSQLRepository) GetList() ([]models.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description, status, created_at, updated_at FROM tasks ORDER BY created_at")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *PostgreSQLRepository) UpdateTaskByID(task models.Task) error {
	query := "UPDATE tasks SET title=$1, description=$2, status=$3, updated_at=now() WHERE id=$4"
	_, err := r.db.Exec(query, task.Title, task.Description, task.Status, task.Id)

	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	return nil
}

func (r *PostgreSQLRepository) DeleteTaskByID(id uint) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
