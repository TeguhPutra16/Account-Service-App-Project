package controllers

import (
	"database/sql"
	"fmt"
	"log"
)

func DeleteAccount(db *sql.DB, read int) error {
	var query = "DELETE FROM users where id=?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {

		return errPrepare
	}

	result, errExec := statement.Exec(read)
	if errExec != nil {

		return errExec
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			log.Fatal("success delete")
		} else {
			fmt.Println("Failed to Delete Account")
		}
	}
	return nil

}
