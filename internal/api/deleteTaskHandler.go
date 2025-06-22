package api

import (
	"net/http"

	"github.com/azoma13/ToDo_List_Practicum/internal/db"
	"github.com/azoma13/ToDo_List_Practicum/internal/service"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := db.DeleteTaskDB(id); err != nil {
		service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	service.SendResponse(w, http.StatusOK, struct{}{})
}
