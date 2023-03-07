package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func PasswdGen(passwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
	if err != nil {
		log.Printf("%v", err)
	}
	return string(hash)
}
