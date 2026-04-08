package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/l10-bhushan/task_tracker_api/internal/model"
	"github.com/l10-bhushan/task_tracker_api/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: s,
	}
}

// Getting all tasks
func (t *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := t.service.GetAllTasks()
	json.NewEncoder(w).Encode(tasks)
}

// Fetching data based on ID
func (t *TaskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	var req model.Task

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid req body", http.StatusBadRequest)
		return
	}

	task := t.service.GetTaskById(req.Id)
	json.NewEncoder(w).Encode(task)
}

// Handler for creating a task
func (t *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req model.Task

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	task := t.service.CreateTask(req.Title, req.Category)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Deleting the task
func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	t.service.DeleteTask(id)
	w.WriteHeader(http.StatusNoContent)
}

// Updating the task
func (t *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var req model.Task

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid req body", http.StatusBadRequest)
		return
	}

	task, result := t.service.UpdateTask(req.Id, req)
	if !result {
		http.Error(w, "Updation failed...", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
