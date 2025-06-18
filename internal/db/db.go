package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/azoma13/ToDo_List_Practicum/configs"
	_ "modernc.org/sqlite"
)

const (
	schema = `
		CREATE TABLE scheduler (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		date CHAR(8) NOT NULL DEFAULT "",
			title VARCHAR NOT NULL DEFAULT "",
			comment TEXT NOT NULL DEFAULT "",
			repeat VARCHAR NOT NULL DEFAULT ""
		);
		CREATE INDEX idx_date ON scheduler (date);
	`
)

var DB *sql.DB

func Init(dbFile string) error {
	envDBFile := configs.ToDoListPort
	if envDBFile != "" {
		dbFile = envDBFile
	}
	_, err := os.Stat(dbFile)

	install := false
	if err != nil {
		install = true
	}

	DB, err = sql.Open("sqlite", "scheduler.db")
	if err != nil {
		return fmt.Errorf("error open database: %v", err)
	}

	if install {
		_, err = DB.Exec(schema)
		if err != nil {
			return fmt.Errorf("error exec sql command schema")
		}
	}

	return nil
}
