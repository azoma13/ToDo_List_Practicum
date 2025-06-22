package api

import (
	"net/http"

	"github.com/azoma13/ToDo_List_Practicum/internal/service"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodGet:
		getTaskByIdHandler(w, r)
	case http.MethodPut:
		updateTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r)
	default:
		service.SendResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "error method not allowed"})
	}
}
