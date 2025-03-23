package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) string {
	log.Println("Hashing password")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("Error while hashing ",err)
	}
	return string(hash)
}

func ComparePassword(hashedPassword, password string) error {
	log.Println("Comparing password")
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

