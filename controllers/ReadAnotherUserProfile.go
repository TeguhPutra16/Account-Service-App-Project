package controllers

import (
	"be13/account-service-app-project/entities"
	"database/sql"
	"fmt"
	"log"
)

func ReadAnotherUserProfile(db *sql.DB, readUser entities.Users) {
	// membuka profile user (selain data credential) dengan id
	result := db.QueryRow("select id, name, gender, address, email, created_at, updated_at from users where id = ?", readUser.Id)
	errScan := result.Scan(&readUser.Id, &readUser.Name, &readUser.Gender, &readUser.Address, &readUser.Email, &readUser.Created_at, &readUser.Updated_at)
	if errScan != nil {
		log.Fatal("User not found, check id", errScan.Error())
	}
	fmt.Printf("\nWELCOME IN YOUR ACCOUNT\nYour Data:\nId: %d\nName: %s\nGender: %s\nAddress: %s\nEmail: %s\nCreated at: %s\nUpdated at: %s\n", readUser.Id, readUser.Name, readUser.Gender, readUser.Address, readUser.Email, readUser.Created_at, readUser.Updated_at)
}
