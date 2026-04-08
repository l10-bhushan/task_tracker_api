package router

import (
	"github.com/gorilla/mux"
	"github.com/l10-bhushan/task_tracker_api/internal/handler"
)

// Setting up the router
func SetupRouter(t *handler.TaskHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/task/create", t.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/task/getAll", t.GetAllTasks).Methods("GET")
	r.HandleFunc("/task/get/{id}", t.GetTaskById).Methods("GET")
	r.HandleFunc("/task/delete/{id}", t.DeleteTask).Methods("DELETE")
	r.HandleFunc("/task/update/", t.UpdateTask).Methods("PUT", "PATCH")

	return r
}
