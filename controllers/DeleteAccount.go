package controllers

import (
	"database/sql"
	"fmt"
	"os"
)

func DeleteAccount(db *sql.DB, read int, yesNo string) error {

	//Konfirmasi delete Account
	if yesNo == "n" {
		return nil // kembali ke menu
	} else if yesNo == "y" {
		/// Proses dilanjutkan
	} else {
		fmt.Println("Submit error")
		os.Exit(1) // keluar sistem
	}

	/// Proses Delete Account
	///Variable read adalah data user id yang diambil dari fungsi login
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
			fmt.Println("Account has been successfully deleted")
		} else {
			fmt.Println("Failed to Delete Account")
		}
	}
	return nil

}
