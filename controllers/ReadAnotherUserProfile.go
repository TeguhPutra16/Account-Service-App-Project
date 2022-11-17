package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
)

func ReadAnotherUserProfile(db *sql.DB, readUser entities.Users) {
	// membuka profile user (selain data credential) dengan id
	result := db.QueryRow("select id, name, gender, address, email, telp_number from users where telp_number = ?", readUser.Telp_number)
	errScan := result.Scan(&readUser.Id, &readUser.Name, &readUser.Gender, &readUser.Address, &readUser.Email, &readUser.Telp_number)
	if errScan != nil {
		log.Fatal("User not found, check id", errScan.Error())
	}
	fmt.Printf("\n==================================================\nUser profile:\nId: %d\nName: %s\nGender: %s\nAddress: %s\nEmail: %s\nTelp Number: %s\n==================================================\n", readUser.Id, readUser.Name, readUser.Gender, readUser.Address, readUser.Email, readUser.Telp_number)
}
