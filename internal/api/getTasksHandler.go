package api

import (
	"net/http"
	"time"

	"github.com/azoma13/ToDo_List_Practicum/configs"
	"github.com/azoma13/ToDo_List_Practicum/internal/db"
	"github.com/azoma13/ToDo_List_Practicum/internal/service"
	"github.com/azoma13/ToDo_List_Practicum/models"
)

func getTasksHandler(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")
	var tasks []*models.Task
	if search == "" {
		var err error
		tasks, err = db.GetTasksDB(configs.Limit)
		if err != nil {
			service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
	} else {
		var err error
		tasks, err = getTasksWithSearch(search, configs.Limit)
		if err != nil {
			service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
	}

	service.SendResponse(w, http.StatusOK, models.TasksResponse{
		Tasks: tasks,
	})
	return
}

func getTasksWithSearch(search string, limit int) ([]*models.Task, error) {
	searchTime, err := time.Parse("02.01.2006", search)
	if err != nil {
		search = "%" + search + "%"
		return db.GetTasksWithSearchDB(search, limit)
	}

	search = searchTime.Format(configs.DateFormat)
	return db.GetTasksWithSearchDateDB(search, limit)
}
