package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

var (
	Obj cryptoInterface
)

func init() {
	Obj = &cryptoObj{}
}

type cryptoInterface interface {
	GenerateHash(password string) (string, error)
	IsPWAndHashPWMatch(password string, hashPass string) bool
}

type cryptoObj struct {
}

func (c *cryptoObj) GenerateHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(passwordHash), err
}

func (c *cryptoObj) IsPWAndHashPWMatch(password string, hashPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if err != nil {
		return false
	}
	return true
}
