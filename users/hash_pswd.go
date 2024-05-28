package users

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPswd(orgPswd string) (hashPswd string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(orgPswd), 14)
	if err != nil {
		log.Fatal(err)
	}

	hashPswd = string(bytes)

	return hashPswd
}
