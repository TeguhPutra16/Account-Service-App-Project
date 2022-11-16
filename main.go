package main

import (
	"be13/project/config"
	"be13/project/controllers"
	"be13/project/entities"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnection := config.Connection()

	for {
		fmt.Println("\nMENU:\n1. ADD ACCOUNT\n2. LOGIN\n0. Exit")
		fmt.Println("\nMasukkan pilihan anda: ")
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			{
				addAccount := entities.Users{}

				fmt.Println("\nEnter your Name:")
				fmt.Scanln(&addAccount.Name)
				fmt.Println("Enter your Email:")
				fmt.Scanln(&addAccount.Email)
				fmt.Println("Insert Password:")
				fmt.Scanln(&addAccount.Password)
				fmt.Println("Enter your Address:")
				fmt.Scanln(&addAccount.Address)
				fmt.Println("Telephone Number:")
				fmt.Scanln(&addAccount.Telp_number)
				fmt.Println("your Balance:")
				fmt.Scanln(&addAccount.Balance)
				fmt.Println("Gender (M/F):")
				fmt.Scanln(&addAccount.Gender)

				err := controllers.Addaccount(dbConnection, addAccount)
				if err != nil {
					log.Fatal("Failed to create Account")
				}
			}
		case 2:
			{
				loginAccount := entities.Users{}
				var pass string

				fmt.Println("\nEnter your Telp number:")
				fmt.Scanln(&loginAccount.Telp_number)
				fmt.Println("Enter your Password:")
				fmt.Scanln(&pass)

				read, errLogin := controllers.LoginAccount(dbConnection, loginAccount, pass)

				if errLogin != nil {
					log.Fatal("Login Failed")
				} else {
					for {
						fmt.Println()
						fmt.Println("Menu: \n1. READ ACCOUNT\n2. UPDATE ACCOUNT\n3. DELETE ACCOUT\n4. TOP UP\n5. TRANSFER\n6. TOP UP HISTORY\n7. TRANSFER HISTORY\n8. READ ANOTHER USER PROFILE\n0. EXIT")
						fmt.Println("\nChoose Menu: ")
						var pilihan int
						fmt.Scanln(&pilihan)
						switch pilihan {
						case 1:
							{

								v := controllers.ReadAccount(dbConnection, read)

								fmt.Printf("\n\n===================\nUser Id: %d\nName: %s\nEmail: %s\nUser Balance: %d\nAddress: %s\nGender: %s\nTelephone:%s\n===================", v.Id, v.Name, v.Email, v.Balance, v.Address, v.Gender, v.Telp_number)

							}
						case 2:
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
						case 3:
							{
								delete := entities.Users{}

								fmt.Println("Enter user id you want to delete:")
								fmt.Scanln(&delete.Id)

								errdelete := controllers.DeleteAccount(dbConnection, delete)
								if errdelete != nil {
									fmt.Println("Delete Failed")
								}
							}
						case 4:
							{
								var topupuserNumber string
								var topupamount int

								fmt.Println("\nEnter your Telp number to Top up:")
								fmt.Scanln(&topupuserNumber)
								fmt.Println("Enter the top up nominal:")
								fmt.Scanln(&topupamount)

								controllers.BalanceTopUp(dbConnection, topupuserNumber, topupamount)
							}
						case 5:
							{
								var nomor, nomor1 string
								var Jum_Tf int

								fmt.Println("\nEnter your Phone number:")
								fmt.Scanln(&nomor)
								fmt.Println("Enter recipient number:")
								fmt.Scanln(&nomor1)
								fmt.Println("Transfer amount:")
								fmt.Scanln(&Jum_Tf)

								controllers.Transfer(dbConnection, nomor, nomor1, Jum_Tf)
							}
						case 6:
							{
								user := entities.Users{}

								fmt.Println("\nEnter User Id to look TopUp History:")
								fmt.Scanln(&user.Id)

								cek := controllers.TopUpHistory(dbConnection, user)
								if len(cek) == 0 {
									log.Fatal("User Id TopUp History not found")
								}
							}
						case 7:
							{
								history := entities.Users{}

								fmt.Println("Enter user id")
								fmt.Scanln(&history.Id)

								cek := controllers.TransferHistory(dbConnection, history)
								if len(cek) == 0 {
									log.Fatal("History not Found")
								}
							}
						case 8:
							{
								readUser := entities.Users{}
								fmt.Println("\nEnter user id:")
								fmt.Scanln(&readUser.Id)

								controllers.ReadAnotherUserProfile(dbConnection, readUser)
							}
						case 0:
							{
								controllers.Exit()

							}

						}

					}

				}

			}
		case 0:
			controllers.Exit()
		}

	}
}
