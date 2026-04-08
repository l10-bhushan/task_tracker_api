package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/l10-bhushan/task_tracker_api/internal/handler"
	"github.com/l10-bhushan/task_tracker_api/internal/repository"
	"github.com/l10-bhushan/task_tracker_api/internal/router"
	"github.com/l10-bhushan/task_tracker_api/internal/service"
)

func main() {
	fmt.Println("Task Tracker API")

	repo := repository.NewTaskRepo()
	service := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(service)

	r := router.SetupRouter(handler)

	log.Println("Serving at PORT :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
