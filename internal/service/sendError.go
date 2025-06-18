package service

import (
	"encoding/json"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, statusServer int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(statusServer)
	json.NewEncoder(w).Encode(map[string]string{
		"error": errorMessage,
	})
}
