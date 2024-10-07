package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "vukivuki"
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Хэш пароля: %s", passHash)
}
