package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"restapi.com/dagem/utils"
)

func Authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "UnAuthorized"})
		return
	}
	UserId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	context.Set("userId", UserId)
	context.Next()
}
