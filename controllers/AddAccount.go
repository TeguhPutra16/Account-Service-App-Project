package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Addaccount(db *sql.DB, insert entities.Users) error {
	// x := Bcript(insert.Password)
	// insert.Password = x
	var query = "insert into users (name,email,password,address,telp_number,balance,gender) values (?,?,?,?,?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		// log.Fatal("error prepare insert", errPrepare.Error())
		return errPrepare
	}

	result, errExec := statement.Exec(insert.Name, insert.Email, insert.Password, insert.Address, insert.Telp_number, insert.Balance, insert.Gender)
	if errExec != nil {
		// log.Fatal("error exec insert", errExec.Error())
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
