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

	fmt.Println("Menu:\n 1.Add Account \n 2.login \n 3.Read Account \n 4.Update Account \n 5.Delete Account \n 6.Top Up \n 7.Transfer")
	fmt.Println("Masukan pilihan anda:")
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

	case 3:

		NoId := entities.Users{}
		fmt.Println("masukkan id user")
		fmt.Scanln(&NoId.Id)

		v := controllers.ReadAccount(dbConnection, NoId)

		fmt.Printf("Id:%d	Name:%s	Email:%s	Password:%s	Address:%s	Balance:%d	Gender:%s	Telephone:%s	date:%s\n", v.Id, v.Name, v.Email, v.Password, v.Address, v.Balance, v.Gender, v.Telp_number, v.Created_at)

	case 4:

	case 5:
		delete := entities.Users{}
		fmt.Println("Masukkan id user yang akan di hapus")
		fmt.Scanln(&delete.Id)

		errdelete := controllers.DeleteAccount(dbConnection, delete)
		if errdelete != nil {
			fmt.Println("Delete Failed")
		}
	case 6:
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

	}

	defer dbConnection.Close()

}
