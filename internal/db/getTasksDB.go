package db

import (
	"database/sql"
	"fmt"

	"github.com/azoma13/ToDo_List_Practicum/models"
)

func GetTasksDB(limit int) ([]*models.Task, error) {
	query := `
		SELECT * FROM scheduler
		ORDER BY date ASC LIMIT ?
	`

	rows, err := DB.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("error query for get tasks")
	}
	defer rows.Close()

	return returnRows(rows)
}

func GetTasksWithSearchDB(search string, limit int) ([]*models.Task, error) {
	query := `SELECT * FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date LIMIT ?`

	rows, err := DB.Query(query, search, search, limit)
	if err != nil {
		return nil, fmt.Errorf("error query for get tasks")
	}
	defer rows.Close()

	return returnRows(rows)
}

func GetTasksWithSearchDateDB(search string, limit int) ([]*models.Task, error) {
	query := `
		SELECT * FROM scheduler 
		WHERE date = ? LIMIT ?
	`

	rows, err := DB.Query(query, search, limit)
	if err != nil {
		return nil, fmt.Errorf("error query for get tasks")
	}
	defer rows.Close()

	return returnRows(rows)
}

func returnRows(rows *sql.Rows) ([]*models.Task, error) {
	tasks := []*models.Task{}
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, fmt.Errorf("error scan rows for get tasks")
		}
		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error cursor rows in func getTaskDB: %w", err)
	}

	return tasks, nil
}
