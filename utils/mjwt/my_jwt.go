package mjwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/muchlist/erru_utils_go/rest_err"
	"os"
)

var (
	Obj jwtServiceInterface
)

func init() {
	Obj = &jwtUtils{}
}

type jwtServiceInterface interface {
	GenerateToken(userEmail string, isAdmin bool) (string, rest_err.APIError)
	ValidateToken(tokenString string) (*jwt.Token, rest_err.APIError)
}

type jwtUtils struct {
}

const secretKey = "SECRET_KEY"

var (
	secret = []byte(os.Getenv(secretKey))
)

func (j *jwtUtils) GenerateToken(userEmail string, isAdmin bool) (string, rest_err.APIError) {
	jwtClaim := jwt.MapClaims{}
	jwtClaim["user_email"] = userEmail
	jwtClaim["is_admin"] = isAdmin
	jwtClaim["exp"] = nil
	jwtClaim["jti"] = nil

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)

	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", rest_err.NewInternalServerError("gagal menandatangani token", err)
	}

	return signedToken, nil
}

/*if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    fmt.Println(claims["foo"], claims["nbf"])
} else {
    fmt.Println(err)
}*/
func (j *jwtUtils) ValidateToken(tokenString string) (*jwt.Token, rest_err.APIError) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, rest_err.NewBadRequestError("token tidak valid")
		}
		return secret, nil
	})

	if err != nil {
		return nil, rest_err.NewBadRequestError("token tidak valid 2")
	}

	return token, nil
}
