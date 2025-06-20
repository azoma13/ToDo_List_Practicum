package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/azoma13/ToDo_List_Practicum/configs"
	"github.com/azoma13/ToDo_List_Practicum/internal/db"
	"github.com/azoma13/ToDo_List_Practicum/internal/service"
	"github.com/azoma13/ToDo_List_Practicum/models"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	id, err := db.AddTaskDB(&task)
	if err != nil {
		service.SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	service.SendResponse(w, http.StatusOK, map[string]int64{"id": id})
}

func checkDate(task *models.Task) error {

	nowTime := time.Now()
	now := nowTime.Format(configs.DateFormat)
	if task.Date == "" {
		task.Date = now
		return nil
	}

	t, err := time.Parse(configs.DateFormat, task.Date)
	if err != nil {
		return fmt.Errorf("error field date is not correct")
	}
	var next string
	if task.Repeat != "" {
		next, err = service.NextDate(nowTime, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if service.AfterNow(t, nowTime) {
		if len(task.Repeat) == 0 {
			task.Date = now
		} else {
			task.Date = next
		}
	}

	return nil
}
