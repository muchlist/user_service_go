package ping_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "PONG!", "time": time.Now()})
}
