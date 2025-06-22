package main

import (
	"log"

	"github.com/azoma13/ToDo_List_Practicum/configs"
	"github.com/azoma13/ToDo_List_Practicum/internal/db"
	"github.com/azoma13/ToDo_List_Practicum/internal/server"
)

func main() {

	err := configs.Environment()
	if err != nil {
		log.Fatalf("Error environment func: %v", err)
	}

	err = db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("error init data base: %v", err)
	}

	err = server.Run()
	if err != nil {
		panic(err)
	}
}
