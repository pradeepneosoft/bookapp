package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass []byte) string {

	hash, err := bcrypt.GenerateFrompassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("failed to hash password")
	}
	return string(hash)
}
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err)
		return false

	}
	return true

}
