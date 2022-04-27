package service

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

const (
	minCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	maxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	defaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

func CryptPassword(password string) []byte {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), defaultCost)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return pwd
}

// Check user password correctness
func CheckPwd(password []byte, cred []byte) bool {
	err := bcrypt.CompareHashAndPassword(password, cred)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
