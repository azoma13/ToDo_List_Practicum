package db

import (
	"fmt"

	"github.com/azoma13/ToDo_List_Practicum/models"
)

func UpdateTaskDB(task *models.Task) error {
	query := `
		UPDATE scheduler 
		SET date = ?, title = ?, comment = ?, repeat = ? 
		WHERE id = ?
	`

	res, err := DB.Exec(query, &task.Date, &task.Title, &task.Comment, &task.Repeat, &task.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}

	return nil
}
