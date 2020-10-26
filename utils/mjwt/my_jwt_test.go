package mjwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJwtUtils_GenerateToken(t *testing.T) {
	email := "muchlis@gmail.com"
	isAdmin := true

	jwt := NewJWTService()
	signedToken, err := jwt.GenerateToken(email, isAdmin)
	println(signedToken)

	assert.Nil(t, err)
	assert.NotEmpty(t, signedToken)
}

func TestJwtUtils_ValidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOm51bGwsImlzX2FkbWluIjp0cnVlLCJqdGkiOm51bGwsInVzZXJfZW1haWwiOiJtdWNobGlzQGdtYWlsLmNvbSJ9.y6UqAuj98zVGSwzps8PMgvdEaeEfvyYbdDbiFaWte9s"
	jwt := NewJWTService()
	tokenValid, err := jwt.ValidateToken(token)

	println(tokenValid)

	assert.Nil(t, err)
	assert.NotEmpty(t, tokenValid)
}
