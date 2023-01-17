package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pwd := `Hola12Amigos`

	bcr, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error is: ", err)
	}
	fmt.Println(pwd)
	fmt.Println(bcr)

	loginPwd := `Hola12Amigos`

	err = bcrypt.CompareHashAndPassword(bcr, []byte(loginPwd))
	if err != nil {
		fmt.Println("Error is: ", err)
		return
	}
	fmt.Println("Password correct! You are logged in.")
}
