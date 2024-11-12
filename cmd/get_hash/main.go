package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "1a2b3cC#"
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Хэш пароля: %s", passHash)
}
