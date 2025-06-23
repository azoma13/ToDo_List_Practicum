package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	DateFormat = "20060102"
	WebDir     = "./web"
	Limit      = 50
)

var (
	ToDoListPort   string
	ToDoListDBFile string
	TODOPassword   string
	JwtKey         = []byte("jwt_secret") // хардкод
)

func Environment() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error load .env: %d", err)
	}

	ToDoListPort = os.Getenv("TODOLIST_PORT")
	ToDoListDBFile = os.Getenv("TODOLIST_DBFILE")
	TODOPassword = os.Getenv("TODO_PASSWORD")
	return nil
}
