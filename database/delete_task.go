package database

import (
	"database/sql"
	//"github.com/foxmeller/go_final_project/models"
)

func DoneTask(db *sql.DB, id int) error {
	_, err := db.Exec(("DELETE FROM scheduler WHERE id = ?"), id)
	return err
}
