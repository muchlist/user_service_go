package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/user_service_go/utils/mjwt"
	"net/http"
	"strings"
	"time"
)

const (
	headerKey = "Authorization"
	bearerKey = "Bearer"
)

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

	if claims.Exp < time.Now().Unix() {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			rest_err.NewAPIError("Token expired", http.StatusUnprocessableEntity, "jwt_expired", nil))
		return
	}

	c.Set(mjwt.CLAIMS, claims)
}