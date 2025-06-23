package api

import (
	"net/http"

	"github.com/azoma13/ToDo_List_Practicum/internal/service"
)

func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", service.Auth(taskHandler))
	http.HandleFunc("/api/tasks", service.Auth(getTasksHandler))
	http.HandleFunc("/api/task/done", service.Auth(doneTaskHandler))
	http.HandleFunc("/api/signin", signinHandler)
}
