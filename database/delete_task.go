package database

import (
	"database/sql"
	"log"

	"github.com/foxmeller/go_final_project/models"
)

func DoneTask(db *sql.DB, id int) (*models.Task, error) {
	_, err := db.Exec(("DELETE FROM scheduler WHERE id = ?"), id)
	if err != nil {
		log.Printf("failed to delete task with ID: %d", id)
		return nil, err
	}

	return &models.Task{}, err
}
