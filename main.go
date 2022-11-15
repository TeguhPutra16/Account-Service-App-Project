package main

import (
	"be13/account-service-app-project/config"
	"be13/account-service-app-project/controllers"
	"be13/account-service-app-project/entities"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnection := config.Connection()

	defer dbConnection.Close()

	fmt.Println("\nMENU:\n1. ADD ACCOUNT\n2. LOGIN\n3. READ ACCOUNT\n4. UPDATE ACCOUNT\n5. DELETE ACCOUT\n6. BALANCE TOP UP\n7. BALANCE TRANSFER\n8. BALANCE TOP UP HISTORY\n9. BALANCE TRANSFER HISTORY\n10. READ ANOTHER USER PROFILE\n0. EXIT")
	fmt.Println("\nMasukkan pilihan anda: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		{

		}
	case 2:
		{
			loginAccount := entities.Users{}
			fmt.Println("\nEnter your Telp number:")
			fmt.Scanln(&loginAccount.Telp_number)
			fmt.Println("Enter your Password:")
			fmt.Scanln(&loginAccount.Password)

			controllers.LoginAccount(dbConnection, loginAccount)
		}
	case 3:
		{

		}
	case 4:
		{
			updateAccount := entities.Users{}
			fmt.Println("\nEnter the user id you want to update:")
			fmt.Scanln(&updateAccount.Id)
			fmt.Println("Enter name update:")
			fmt.Scanln(&updateAccount.Name)
			fmt.Println("Enter gender update (M/F):")
			fmt.Scanln(&updateAccount.Gender)
			fmt.Println("Enter address update:")
			fmt.Scanln(&updateAccount.Address)
			fmt.Println("Enter email update:")
			fmt.Scanln(&updateAccount.Email)
			fmt.Println("Enter telp number update:")
			fmt.Scanln(&updateAccount.Telp_number)
			fmt.Println("Enter password update:")
			fmt.Scanln(&updateAccount.Password)

			controllers.UpdateAccount(dbConnection, updateAccount)
		}
	case 5:
		{

		}
	case 6:
		{

		}
	case 7:
		{

		}
	case 8:
		{

		}
	case 9:
		{

		}
	case 10:
		{
			readUser := entities.Users{}
			fmt.Println("\nEnter id:")
			fmt.Scanln(&readUser.Id)

			controllers.ReadAnotherUserProfile(dbConnection, readUser)
		}
	case 0:
		{

		}
	}
}
