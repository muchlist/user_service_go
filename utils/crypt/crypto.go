package crypt

import (
	"github.com/muchlist/erru_utils_go/logger"
	"github.com/muchlist/erru_utils_go/rest_err"
	"golang.org/x/crypto/bcrypt"
)

var (
	Obj cryptoInterface
)

func init() {
	Obj = &cryptoObj{}
}

type cryptoInterface interface {
	GenerateHash(password string) (string, rest_err.APIError)
	IsPWAndHashPWMatch(password string, hashPass string) bool
}

type cryptoObj struct {
}

func (c *cryptoObj) GenerateHash(password string) (string, rest_err.APIError) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		logger.Error("Error pada kriptograpi (GenerateHash)", err)
		restErr := rest_err.NewInternalServerError("Crypto error", err)
		return "", restErr
	}
	return string(passwordHash), nil
}

func (c *cryptoObj) IsPWAndHashPWMatch(password string, hashPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if err != nil {
		return false
	}
	return true
}
