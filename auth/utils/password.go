package utils

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[utils (HashPassword)] \n ")
		return "", errors.New(fmt.Sprintf("error trying to hash password ➜ %v", err))
	}

	return string(hash), nil
}

func VerifyPassword(loginPass, storedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(storedPass), []byte(loginPass))
}
