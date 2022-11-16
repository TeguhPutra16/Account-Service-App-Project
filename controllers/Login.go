package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
)

func LoginAccount(db *sql.DB, loginAccount entities.Users) {
	// login untuk 1 account
	result := db.QueryRow("select id, name, gender, address, email, telp_number, password, balance from users where telp_number = ? and password = ?", loginAccount.Telp_number, loginAccount.Password)
	errScan := result.Scan(&loginAccount.Id, &loginAccount.Name, &loginAccount.Gender, &loginAccount.Address, &loginAccount.Email, &loginAccount.Telp_number, &loginAccount.Password, &loginAccount.Balance)
	if errScan != nil {
		log.Fatal("User not found, check your Telp number and Password again", errScan.Error())
	}
	fmt.Printf("\nWELCOME IN YOUR ACCOUNT\nYour Data:\nId: %d\nName: %s\nGender: %s\nAddress: %s\nEmail: %s\nTelp Number: %s\nPassword: %s\nBalance: %d\n", loginAccount.Id, loginAccount.Name, loginAccount.Gender, loginAccount.Address, loginAccount.Email, loginAccount.Telp_number, loginAccount.Password, loginAccount.Balance)
}
