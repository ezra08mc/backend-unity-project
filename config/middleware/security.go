package middleware

import (
	"net/http"
	"strings"

	"github.com/ezra08mc/backend-unity-project/config/pkg/errs"
	"github.com/ezra08mc/backend-unity-project/config/pkg/token"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Authorization header required",
				"code":    "MISSING_TOKEN",
			})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Invalid authorization format. Use: Bearer <token>",
				"code":    "INVALID_TOKEN_FORMAT",
			})
			ctx.Abort()
			return
		}

		tokenString := parts[1]

		userData, err := token.ValidateAccessToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Invalid or expired token",
				"code":    "INVALID_TOKEN",
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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errs.Unauthorized("user not authenticated"))
			return
		}

		if role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, errs.Forbidden("admin access required"))
			return
		}

		ctx.Next()
	}
}
