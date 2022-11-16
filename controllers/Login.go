package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func LoginAccount(db *sql.DB, loginAccount entities.Users, pass string) {
	// login untuk 1 account
	result1 := db.QueryRow("select password from users where telp_number = ?", loginAccount.Telp_number) //dapat data password
	errScan1 := result1.Scan(&loginAccount.Password)
	if errScan1 != nil {
		log.Fatal("User not found, check your Telp number and Password again", errScan1.Error())
	}

	hashed_Pass := loginAccount.Password

	z := []byte(pass)
	err := bcrypt.CompareHashAndPassword([]byte(hashed_Pass), z)
	if err != nil {
		log.Fatal("Password salah")
	}

	result := db.QueryRow("select id, name, gender, address, email, telp_number, password, balance from users where telp_number = ? and password = ?", loginAccount.Telp_number, loginAccount.Password)
	errScan := result.Scan(&loginAccount.Id, &loginAccount.Name, &loginAccount.Gender, &loginAccount.Address, &loginAccount.Email, &loginAccount.Telp_number, &loginAccount.Password, &loginAccount.Balance)
	if errScan != nil {
		log.Fatal("User not found, check your Telp number and Password again", errScan.Error())
	}
	fmt.Printf("\nWELCOME IN YOUR ACCOUNT\nYour Data:\nId: %d\nName: %s\nGender: %s\nAddress: %s\nEmail: %s\nTelp Number: %s\nPassword: %s\nBalance: %d\n", loginAccount.Id, loginAccount.Name, loginAccount.Gender, loginAccount.Address, loginAccount.Email, loginAccount.Telp_number, loginAccount.Password, loginAccount.Balance)
}
