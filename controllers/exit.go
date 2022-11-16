package controllers

import (
	"fmt"
	"os"
)

func Exit() {
	fmt.Println()

	fmt.Println("Terimakasih telah bertransaksi :)")

	os.Exit(0)

}
