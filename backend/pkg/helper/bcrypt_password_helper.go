package helper

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHelper struct{}

func (bph BcryptPasswordHelper) Hash(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	return string(hashedPasswordBytes), err
}

func (bph BcryptPasswordHelper) IsMatch(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(password))
	return err == nil
}

func NewBcryptPasswordHelper() *BcryptPasswordHelper {
	return &BcryptPasswordHelper{}
}
