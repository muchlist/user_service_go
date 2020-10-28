package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/user_service_go/utils/mjwt"
	"strings"
)

const (
	headerKey = "Authorization"
	bearerKey = "Bearer"
)

//AuthMiddleware memvalidasi token JWT, mengembalikan claims berupa pointer mjwt.CustomClaims
func AuthMiddleware(c *gin.Context) {

	authHeader := c.GetHeader(headerKey)
	if !strings.Contains(authHeader, bearerKey) {
		apiErr := rest_err.NewUnauthorizedError("Unauthorized")
		c.AbortWithStatusJSON(apiErr.Status(), apiErr)
		return
	}

	tokenString := strings.Split(authHeader, " ")
	if len(tokenString) != 2 {
		apiErr := rest_err.NewUnauthorizedError("Unauthorized")
		c.AbortWithStatusJSON(apiErr.Status(), apiErr)
		return
	}

	token, err := mjwt.Obj.ValidateToken(tokenString[1])
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	claims, err := mjwt.Obj.ReadToken(token)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.Set(mjwt.CLAIMS, claims)
}

//AuthAdminMiddleware memvalidasi token JWT, mengembalikan claims berupa pointer mjwt.CustomClaims
//perbedaannya dengan AuthMidlleware adalah ini mengharuskan pengakses berstatus is_admin true
func AuthAdminMiddleware(c *gin.Context) {

	authHeader := c.GetHeader(headerKey)
	if !strings.Contains(authHeader, bearerKey) {
		apiErr := rest_err.NewUnauthorizedError("Unauthorized")
		c.AbortWithStatusJSON(apiErr.Status(), apiErr)
		return
	}

	tokenString := strings.Split(authHeader, " ")
	if len(tokenString) != 2 {
		apiErr := rest_err.NewUnauthorizedError("Unauthorized")
		c.AbortWithStatusJSON(apiErr.Status(), apiErr)
		return
	}

	token, err := mjwt.Obj.ValidateToken(tokenString[1])
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	claims, err := mjwt.Obj.ReadToken(token)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	if !claims.IsAdmin {
		apiErr := rest_err.NewUnauthorizedError("Unauthorized, memerlukan hak akses admin")
		c.AbortWithStatusJSON(apiErr.Status(), apiErr)
		return
	}

	c.Set(mjwt.CLAIMS, claims)
}
