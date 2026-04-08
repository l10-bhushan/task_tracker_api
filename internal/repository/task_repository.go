package repository

import "github.com/l10-bhushan/task_tracker_api/internal/model"

type TaskRepository interface {
	GetAllTasks() []model.Task
	GetTaskById(id string) model.Task
	CreateTask(task model.Task) model.Task
	UpdateTask(id string, task model.Task) model.Task
	DeleteTask(id string)
}
