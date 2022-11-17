package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func DeleteAccount(db *sql.DB, read int, yesNo string) error {

	if yesNo == "n" {
		return nil
	} else if yesNo == "y" {

	} else {
		fmt.Println("Submit error")
		os.Exit(1)
	}
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
			log.Fatal("Account has been successfully deleted")
		} else {
			log.Fatal("Failed to Delete Account")
		}
	}
	return nil

}
