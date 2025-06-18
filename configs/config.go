package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	DateFormat = "20060102"
	WebDir     = "./web"
)

var (
	ToDoListPort   string
	ToDoListDBFile string
)

func Environment() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error load .env: %d", err)
	}

	ToDoListPort = os.Getenv("TODOLIST_PORT")
	ToDoListDBFile = os.Getenv("TODOLIST_DBFILE")
	return nil
}
