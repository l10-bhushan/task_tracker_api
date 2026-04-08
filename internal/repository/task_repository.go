package repository

import "github.com/l10-bhushan/task_tracker_api/internal/model"

// We are creating this interface so that our database would follow these
// Exact rules
// All these methods that you see below will be used to define task service
type TaskRepository interface {
	GetAllTasks() []model.Task
	GetTaskById(id string) (model.Task, bool)
	CreateTask(task model.Task) model.Task
	UpdateTask(id string, task model.Task) (model.Task, bool)
	DeleteTask(id string)
}

// All these methods that you see below are for the in-memory db
// Creating a in memory database that will store our tasks
// Here, as InMemoryTaskRepository implements all the interface methods
// That means InMemoryTaskRepository automatically implements TaskRepository interface
type InMemoryTaskRepository struct {
	data map[string]model.Task
}

// Description: This function would create a db instance and return us a pointer
func NewTaskRepo() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		data: make(map[string]model.Task),
	}
}

// Description: Returns all the tasks in our db
func (r *InMemoryTaskRepository) GetAllTasks() []model.Task {
	var result []model.Task
	for _, value := range r.data {
		result = append(result, value)
	}
	return result
}

// Description: Returns us task for a particular ID
func (r *InMemoryTaskRepository) GetTaskById(id string) (model.Task, bool) {
	task, ok := r.data[id]
	return task, ok
}

// Description: Creates a new task in our db
func (r *InMemoryTaskRepository) CreateTask(task model.Task) model.Task {
	r.data[task.Id] = task
	return task
}

// Description: Deletes a task from our db
func (r *InMemoryTaskRepository) DeleteTask(id string) {
	delete(r.data, id)
}

// Description: Updates a task from our db
func (r *InMemoryTaskRepository) UpdateTask(id string, task model.Task) (model.Task, bool) {
	_, ok := r.data[id]
	if !ok {
		return model.Task{}, false
	}
	r.data[id] = task
	return task, true
}
