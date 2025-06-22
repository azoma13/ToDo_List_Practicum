package db

import "fmt"

func UpdateDateTaskDB(date string, id string) error {
	exec := `
		UPDATE scheduler 
		SET date = ? 
		WHERE id = ?
	`

	_, err := DB.Exec(exec, date, id)
	if err != nil {
		return fmt.Errorf("error execute in func UpdateTask: %w", err)
	}

	return nil
}
