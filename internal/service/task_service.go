package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/l10-bhushan/task_tracker_api/internal/model"
	"github.com/l10-bhushan/task_tracker_api/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: r,
	}
}

func (t *TaskService) CreateTask(title string, category string) model.Task {
	id := uuid.New().String()
	task := model.Task{
		Id:                  id,
		Title:               title,
		Status:              false,
		Timestamp:           time.Now().GoString(),
		CompletionTimestamp: "",
		Category:            category,
	}
	return t.repo.CreateTask(task)
}

func (t *TaskService) GetAllTasks() []model.Task {
	return t.repo.GetAllTasks()
}

func (t *TaskService) GetTaskById(id string) model.Task {
	// Todo : implement this correctly
	// result, err := t.repo.GetTaskById(id)
	// if err
	value, _ := t.repo.GetTaskById(id)
	return value
}
