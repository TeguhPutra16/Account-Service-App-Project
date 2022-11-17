package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Addaccount(db *sql.DB, insert entities.Users) error {

	///Password yang diinput akan dienkripsi dengan package Bcrypt
	x := Bcript(insert.Password)
	insert.Password = x

	///Proses insert data ke tabel Users//////////////////////////////////////////////////////////////////////////

	var query = "insert into users (name,email,password,address,telp_number,balance,gender) values (?,?,?,?,?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {

		return errPrepare
	}
	result, errExec := statement.Exec(insert.Name, insert.Email, insert.Password, insert.Address, insert.Telp_number, insert.Balance, insert.Gender)
	if errExec != nil {

		return errExec
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Account has been Successfully Added")
		} else {
			fmt.Println("Failed to Add Account")
		}
	}
	return nil
}

func Bcript(y string) string {
	password := []byte(y)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(hashedPassword))

	// // Comparing the password with the hash
	// err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	// fmt.Println(err) // nil means it is a match

	return string(hashedPassword)

}
