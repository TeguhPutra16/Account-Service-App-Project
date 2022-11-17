package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func LoginAccount(db *sql.DB, loginAccount entities.Users, pass string) (int, error) {
	// mengambil data passwor yang terenkripsi dari database
	result1 := db.QueryRow("select password from users where telp_number = ?", loginAccount.Telp_number) //dapat data password
	errScan1 := result1.Scan(&loginAccount.Password)
	if errScan1 != nil {
		fmt.Println("User not found, check your Telp number and Password again ")
		os.Exit(1)
	}
	///Proses pengecekan apakah password yang sudah di enkripsi
	///jika pass yang diinput sudah cocok dengan password yang telah terenkripsi maka proses login akan dilanjukan
	hashed_Pass := loginAccount.Password

	z := []byte(pass)
	err := bcrypt.CompareHashAndPassword([]byte(hashed_Pass), z)
	if err != nil {
		return 0, err
	}

	///setelah pencocokan password selesai maka dilanjukan dengan pengambilan data-data user
	result := db.QueryRow("select id, name, gender, address, email, telp_number, password, balance from users where telp_number = ? and password = ?", loginAccount.Telp_number, loginAccount.Password)
	errScan := result.Scan(&loginAccount.Id, &loginAccount.Name, &loginAccount.Gender, &loginAccount.Address, &loginAccount.Email, &loginAccount.Telp_number, &loginAccount.Password, &loginAccount.Balance)
	if errScan != nil {
		log.Fatal("User not found, check your Telp number and Password again")
		os.Exit(2)
	}
	fmt.Printf("\n==================================================\nWelcome to Your Account %s\n==================================================", loginAccount.Name)
	return loginAccount.Id, nil
}
