package db

import "fmt"

func DeleteTaskDB(id string) error {

	res, err := DB.Exec("DELETE FROM scheduler WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error execute in func DeleteTask: %w", err)
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
