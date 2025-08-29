package middlewares

import (
	"net/http"
	"weblog/authentication"

	"github.com/gin-gonic/gin"
)

func ChekAddUserTokenPermit(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	userId, err := authentication.CheckAuthenJWTtoken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authenticated"})
		return
	}

	ctx.Set("userid", userId)

	ctx.Next()
}
