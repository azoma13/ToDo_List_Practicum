package api

import "net/http"

func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", getTasksHandler)
	http.HandleFunc("/api/task/done", doneTaskHandler)
}
