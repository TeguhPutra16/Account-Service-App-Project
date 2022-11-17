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
		fmt.Println("\nEnter your choice: ")
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
				////Data User id saat login ditampung di variabel read
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

								fmt.Printf("\n\n==================================================\nUser Id: %d\nName: %s\nEmail: %s\nGender: %s\nAddress: %s\nUser Balance: %d\nTelephone: %s\nCreated at: %s\nUpdated at: %s\n==================================================\n\n", v.Id, v.Name, v.Email, v.Gender, v.Address, v.Balance, v.Telp_number, v.Created_at, v.Updated_at)

							}
						case 2:
							{
								updateAccount := entities.Users{}

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

								controllers.UpdateAccount(dbConnection, updateAccount, read)
							}
						case 3:
							{
								var yesNo string
								fmt.Println("Are you sure to delete account? (y/n)")
								fmt.Scanln(&yesNo)
								errdelete := controllers.DeleteAccount(dbConnection, read, yesNo)
								if errdelete != nil {
									fmt.Println("Delete Failed")
								}
							}
						case 4:
							{
								var topupamount int

								fmt.Println("Enter the top up nominal:")
								fmt.Scanln(&topupamount)

								controllers.BalanceTopUp(dbConnection, read, topupamount)
							}
						case 5:
							{
								var nomor1, nomor2 string
								var Jum_Tf int

								fmt.Println("Confirm your phone number:")
								fmt.Scanln(&nomor2)
								fmt.Println("Enter recipient number:")
								fmt.Scanln(&nomor1)
								fmt.Println("Transfer amount:")
								fmt.Scanln(&Jum_Tf)

								controllers.Transfer(dbConnection, read, nomor1, nomor2, Jum_Tf)
							}
						case 6:
							{
								cek := controllers.TopUpHistory(dbConnection, read)
								if len(cek) == 0 {
									log.Fatal("History not Found")
								}
							}
						case 7:
							{
								cek := controllers.TransferHistory(dbConnection, read)
								if len(cek) == 0 {
									log.Fatal("History not Found")
								}
							}
						case 8:
							{
								readUser := entities.Users{}
								fmt.Println("\nEnter User Telp Number:")
								fmt.Scanln(&readUser.Telp_number)

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
			{
				controllers.Exit()
			}
		}
	}
}
