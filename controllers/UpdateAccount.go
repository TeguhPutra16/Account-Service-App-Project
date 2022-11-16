package controllers

import (
	"be13/account-service-app-project/entities"
	"database/sql"
	"fmt"
	"log"
)

func UpdateAccount(db *sql.DB, updateAccount entities.Users) {
	// update account berdasarkan id yg ingin di update
	var query = "update users set name = ?, gender = ?, address = ?, email = ?, telp_number = ?, password = ? where id = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare update", errPrepare.Error())
	}

	result, errExec := statement.Exec(updateAccount.Name, updateAccount.Gender, updateAccount.Address, updateAccount.Email, updateAccount.Telp_number, updateAccount.Password, updateAccount.Id)
	if errExec != nil {
		log.Fatal("error exec update", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("update success")
		} else {
			fmt.Println("update failed")
		}
	}
}
