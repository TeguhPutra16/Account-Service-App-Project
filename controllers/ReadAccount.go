package controllers

import (
	"be13/project/entities"
	"database/sql"
	"log"
)

func ReadAccount(db *sql.DB, Id entities.Users) entities.Users {
	result := db.QueryRow("SELECT id,name,email,gender,address,telp_number,balance,password FROM users where id=?", Id.Id)

	var userrow entities.Users
	errScan := result.Scan(&userrow.Id, &userrow.Name, &userrow.Email, &userrow.Gender, &userrow.Address, &userrow.Telp_number, &userrow.Balance, &userrow.Password)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			log.Fatal("Id tdk ada")
		} else {
			log.Fatal("eror scan", errScan.Error())
		}

	}
	// fmt.Printf("id:%s  nama:%s  email: %s Domisili:%s\n", userrow.Id, userrow.Nama, userrow.Email, userrow.Domisili)
	return userrow

}
