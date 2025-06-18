package api

import (
	"net/http"
	"time"

	"github.com/azoma13/ToDo_List_Practicum/configs"
	"github.com/azoma13/ToDo_List_Practicum/internal/service"
)

func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		service.SendErrorResponse(w, http.StatusMethodNotAllowed, "invalid request method")
		return
	}

	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	nowTime, err := time.Parse(configs.DateFormat, now)
	if err != nil {
		service.SendErrorResponse(w, http.StatusBadRequest, "incorrect time format")
		return
	}
	nextDate, err := service.NextDate(nowTime, date, repeat)
	if err != nil {
		service.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = w.Write([]byte(nextDate))
	if err != nil {
		service.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
}
