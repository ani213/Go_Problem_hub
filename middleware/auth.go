package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ani213/Problemhub_backend/common"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenstr := ctx.GetHeader("Authorization")
		mainToken := strings.Split(tokenstr, " ")
		if len(mainToken) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized access",
			})
			return
		}
		secret := os.Getenv("TOKEN_KEY")
		claims, err := common.VerifyToken(mainToken[1], secret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized access",
			})
			return
		}

		ctx.Set("user", claims)
		fmt.Println(mainToken[1])
		ctx.Next()
	}
}
