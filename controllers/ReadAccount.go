package controllers

import (
	"be13/project/entities"
	"database/sql"
	"log"
)

func ReadAccount(db *sql.DB, Id int) entities.Users {

	result := db.QueryRow("SELECT id,name,email,gender,address,telp_number,balance,password,created_at FROM users where id=?", Id)

	var userrow entities.Users
	errScan := result.Scan(&userrow.Id, &userrow.Name, &userrow.Email, &userrow.Gender, &userrow.Address, &userrow.Telp_number, &userrow.Balance, &userrow.Password, &userrow.Created_at)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			log.Fatal("Id tdk ada")
		} else {
			log.Fatal("eror scan", errScan.Error())
		}

	}

	return userrow

}
