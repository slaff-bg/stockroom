package facilities

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// PasswdGen generates a salted hash password.
func PasswdGen(passwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
	if err != nil {
		log.Printf("%v", err)
	}
	return string(hash)
}
