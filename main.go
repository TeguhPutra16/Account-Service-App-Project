package main

import (
	"be13/project/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbConnection := config.Connection()

	fmt.Println("Menu:")
	fmt.Println("Masukan pilihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:

	case 2:

	case 3:

	case 4:

	case 5:

	}

	defer dbConnection.Close()

}
