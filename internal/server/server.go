package server

import (
	"log"
	"net/http"

	"github.com/azoma13/ToDo_List_Practicum/configs"
	"github.com/azoma13/ToDo_List_Practicum/internal/api"
)

func Run() error {
	port := configs.ToDoListPort
	if port == "" {
		port = "7540"
	}
	port = ":" + port

	fileServer := http.FileServer(http.Dir(configs.WebDir))
	http.Handle("/", fileServer)
	api.Init()
	log.Println("application running on port:" + port)
	return http.ListenAndServe(port, nil)
}
