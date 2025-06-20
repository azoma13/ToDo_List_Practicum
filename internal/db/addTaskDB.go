package db

import "github.com/azoma13/ToDo_List_Practicum/models"

func AddTaskDB(task *models.Task) (int64, error) {

	var id int64

	query := `
		INSERT INTO scheduler
			(date, title, comment, repeat)
			VALUES ($1, ?, ?, ?)
	`

	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}

	return id, err
}
