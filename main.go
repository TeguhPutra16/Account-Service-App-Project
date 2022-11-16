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

	defer dbConnection.Close()

	fmt.Println("\nMENU:\n1. ADD ACCOUNT\n2. LOGIN\n3. READ ACCOUNT\n4. UPDATE ACCOUNT\n5. DELETE ACCOUT\n6. BALANCE TOP UP\n7. BALANCE TRANSFER\n8. BALANCE TOP UP HISTORY\n9. BALANCE TRANSFER HISTORY\n10. READ ANOTHER USER PROFILE\n0. EXIT")
	fmt.Println("\nMasukkan pilihan anda: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		addAccount := entities.Users{}

		fmt.Println("Masukan nama")
		fmt.Scanln(&addAccount.Name)
		fmt.Println("Masukan Email")
		fmt.Scanln(&addAccount.Email)
		fmt.Println("Masukan Password")
		fmt.Scanln(&addAccount.Password)
		fmt.Println("Masukan Alamat")
		fmt.Scanln(&addAccount.Address)
		fmt.Println("Masukan No.Telepon")
		fmt.Scanln(&addAccount.Telp_number)
		fmt.Println("Masukan Balance")
		fmt.Scanln(&addAccount.Balance)
		fmt.Println("Jenis Kelamin")
		fmt.Scanln(&addAccount.Gender)

		err := controllers.Addaccount(dbConnection, addAccount)
		if err != nil {
			log.Fatal("Failed to create Account")

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

			NoId := entities.Users{}
			fmt.Println("masukkan id user")
			fmt.Scanln(&NoId.Id)

			v := controllers.ReadAccount(dbConnection, NoId)

			fmt.Printf("Id:%d	Name:%s	Email:%s	Password:%s	Address:%s	Balance:%d	Gender:%s	Telephone:%s	date:%s\n", v.Id, v.Name, v.Email, v.Password, v.Address, v.Balance, v.Gender, v.Telp_number, v.Created_at)

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
		delete := entities.Users{}
		fmt.Println("Masukkan id user yang akan di hapus")
		fmt.Scanln(&delete.Id)

		errdelete := controllers.DeleteAccount(dbConnection, delete)
		if errdelete != nil {
			fmt.Println("Delete Failed")
		}

	case 6:
		{
			var topupuserNumber string
			var topupamount int
			fmt.Println("Enter your Telp number to Top up:")
			fmt.Scanln(&topupuserNumber)
			fmt.Println("Enter the top up nominal:")
			fmt.Scanln(&topupamount)

			controllers.BalanceTopUp(dbConnection, topupuserNumber, topupamount)

		}
	case 7:

		var nomor, nomor1 string
		var Jum_Tf int

		fmt.Println("masukkan nomor anda")
		fmt.Scanln(&nomor)
		fmt.Println("masukkan nomor penerima")
		fmt.Scanln(&nomor1)
		fmt.Println("masukkan nominal Transfer")
		fmt.Scanln(&Jum_Tf)

		controllers.Transfer(dbConnection, nomor, nomor1, Jum_Tf)
	case 8:
		{
			user := entities.Users{}
			fmt.Println("\nEnter Id to look TopUp History:")
			fmt.Scanln(&user.Id)

			controllers.TopUpHistory(dbConnection, user)
		}
	case 9:
		{
			history := entities.Users{}
			fmt.Println("Masukan users id")
			fmt.Scanln(&history.Id)
			cek := controllers.TransferHistory(dbConnection, history)
			if len(cek) == 0 {
				log.Fatal("History not Found")

			}

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
