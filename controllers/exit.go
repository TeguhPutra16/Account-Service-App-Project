package controllers

import (
	"fmt"
	"os"
)

func Exit() {
	fmt.Println()

	fmt.Println("Thank You For Transacting :)")

	os.Exit(0)

}
