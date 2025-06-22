package api

import (
	"net/http"

	"github.com/azoma13/ToDo_List_Practicum/internal/db"
	"github.com/azoma13/ToDo_List_Practicum/internal/service"
)

func getTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	task, err := db.GetTaskByIdDB(id)
	if err != nil {
		service.SendResponse(w, http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
		return
	}

	service.SendResponse(w, http.StatusOK, task)
}
