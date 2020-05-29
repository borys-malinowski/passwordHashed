package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := flag.String("password", "", "Write your password here")
	cost := flag.Int("cost", 10, "Write your cost value here")
	flag.Parse()
	passwodValue := *password
	costValue := *cost
	passwordInBytes := []byte(passwodValue)
	hashedPasswordInBytes, err := bcrypt.GenerateFromPassword(passwordInBytes, costValue)
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword := string(hashedPasswordInBytes)

	fmt.Println(hashedPassword)
}
