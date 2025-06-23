package api

import (
	"net/http"
	"time"

	"github.com/azoma13/ToDo_List_Practicum/internal/db"
	"github.com/azoma13/ToDo_List_Practicum/internal/service"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	task, err := db.GetTaskByIdDB(id)
	if err != nil {
		service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if task.Repeat == "" {

		err := db.DeleteTaskDB(id)
		if err != nil {
			service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

	} else {

		date, err := service.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		err = db.UpdateDateTaskDB(date, id)
		if err != nil {
			service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
	}

	service.SendResponse(w, http.StatusOK, struct{}{})
}
