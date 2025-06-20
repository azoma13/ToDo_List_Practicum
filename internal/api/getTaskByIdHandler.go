package api

import (
	"net/http"
	"strconv"

	"github.com/azoma13/ToDo_List_Practicum/internal/db"
	"github.com/azoma13/ToDo_List_Practicum/internal/service"
)

func getTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		service.SendResponse(w, http.StatusBadRequest, map[string]string{
			"error": "error atoi id",
		})
		return
	}

	task, err := db.GetTaskByIdDB(idInt)
	if err != nil {
		service.SendResponse(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	service.SendResponse(w, http.StatusOK, task)
}
