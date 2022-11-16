package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
)

func DeleteAccount(db *sql.DB, delete entities.Users) error {
	var query = "DELETE FROM users where id=?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {

		return errPrepare
	}

	result, errExec := statement.Exec(delete.Id)
	if errExec != nil {

		return errExec
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Account has been Successfully Deleted")
		} else {
			fmt.Println("Failed to Delete Account")
		}
	}
	return nil

}
