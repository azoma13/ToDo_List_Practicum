package service

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendResponse(w http.ResponseWriter, statusServer int, message any) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(statusServer)
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Println(err)
	}
}
