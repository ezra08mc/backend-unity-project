package middleware

import (
	"net/http"
	"strings"

	"github.com/ezra08mc/backend-unity-project/config/pkg/token"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userData, err := token.ValidateAccessToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", userData.ID)
		ctx.Set("user_email", userData.Email)
		ctx.Set("user_role", userData.Role)

		ctx.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("user_role")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			ctx.Abort()
			return
		}

		if role != "admin" {
			ctx.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}