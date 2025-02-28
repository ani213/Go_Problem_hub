package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ani213/Problemhub_backend/controller"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenstr := ctx.GetHeader("Authorization")
		mainToken := strings.Split(tokenstr, " ")
		if len(mainToken) < 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized access",
			})
		}
		user := controller.User{
			ID:    "1",
			Role:  "admin",
			Email: "ani@gmail.com",
		}
		ctx.Set("user", user)
		fmt.Println(mainToken[1])
		ctx.Next()
	}
}
