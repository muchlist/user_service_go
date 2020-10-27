package mjwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJwtUtils_GenerateToken(t *testing.T) {
	c := CustomClaim{
		Identity:  "muchlis@gmail.com",
		IsAdmin:   true,
		Jti:       "",
		TimeExtra: 12,
	}

	signedToken, err := Obj.GenerateToken(c)
	println(signedToken)

	assert.Nil(t, err)
	assert.NotEmpty(t, signedToken)
}

func TestJwtUtils_ValidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDM4MDcyMzEsImlkZW50aXR5IjoibXVjaGxpc0BnbWFpbC5jb20iLCJpc19hZG1pbiI6dHJ1ZSwianRpIjoiIn0.dzKZdhPFtF-YC6uh5JZqBv7mhBjGTz1_rgIP-sRbYrU"
	tokenValid, err := Obj.ValidateToken(token)

	println(tokenValid)

	assert.Nil(t, err)
	assert.NotEmpty(t, tokenValid)
}

func TestJwtUtils_ReadToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDM4MDcyMzEsImlkZW50aXR5IjoibXVjaGxpc0BnbWFpbC5jb20iLCJpc19hZG1pbiI6dHJ1ZSwianRpIjoiIn0.dzKZdhPFtF-YC6uh5JZqBv7mhBjGTz1_rgIP-sRbYrU"
	tokenValid, err := Obj.ValidateToken(token)
	claims, err2 := Obj.ReadToken(tokenValid)

	assert.Nil(t, err)
	assert.Nil(t, err2)

	assert.Equal(t, "muchlis@gmail.com", claims.Identity)
	assert.Equal(t, true, claims.IsAdmin)
	assert.Equal(t, int64(1603807231), claims.Exp)
	assert.Equal(t, "", claims.Jti)
}
