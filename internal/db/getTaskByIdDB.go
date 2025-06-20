package db

import (
	"fmt"

	"github.com/azoma13/ToDo_List_Practicum/models"
)

func GetTaskByIdDB(id int) (*models.Task, error) {
	query := `
		SELECT * FROM scheduler
		WHERE id = ? 
	`
	task := &models.Task{}
	err := DB.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return &models.Task{}, fmt.Errorf("error select task by id")
	}

	return task, nil
}
