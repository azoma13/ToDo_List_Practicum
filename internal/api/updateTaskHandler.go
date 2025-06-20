package api

import (
	"encoding/json"
	"net/http"

	"github.com/azoma13/ToDo_List_Practicum/internal/db"
	"github.com/azoma13/ToDo_List_Practicum/internal/service"
	"github.com/azoma13/ToDo_List_Practicum/models"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": "error decoding body request"})
		return
	}

	if task.Title == "" {
		service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": "error field title is empty"})
		return
	}

	err = checkDate(&task)
	if err != nil {
		service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err = db.UpdateTaskDB(&task)
	if err != nil {
		service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	service.SendResponse(w, http.StatusOK, struct{}{})
}
