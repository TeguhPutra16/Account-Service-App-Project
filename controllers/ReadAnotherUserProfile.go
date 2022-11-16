package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
)

func ReadAnotherUserProfile(db *sql.DB, readUser entities.Users) {
	// membuka profile user (selain data credential) dengan id
	result := db.QueryRow("select id, name, gender, address, email from users where id = ?", readUser.Id)
	errScan := result.Scan(&readUser.Id, &readUser.Name, &readUser.Gender, &readUser.Address, &readUser.Email)
	if errScan != nil {
		log.Fatal("User not found, check id", errScan.Error())
	}
	fmt.Printf("\nWELCOME IN YOUR ACCOUNT\nYour Data:\nId: %d\nName: %s\nGender: %s\nAddress: %s\nEmail: %s\n", readUser.Id, readUser.Name, readUser.Gender, readUser.Address, readUser.Email)
}
