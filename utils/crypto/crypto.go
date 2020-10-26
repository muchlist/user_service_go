package crypto

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(passwordHash), err
}

func IsPWAndHashPWMatch(password string, hashPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if err != nil {
		return false
	}
	return true
}
