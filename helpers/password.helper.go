package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// accomodate function for hashing

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	fmt.Println(password, string(bytes))
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(hash, password)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, nil
}
