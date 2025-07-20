package middleware

import (
	"net/http"
	"testMod/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(config config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid JWT token",
			})
			ctx.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Database.SecretKey), nil
		})
		if err!=nil{
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}
		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid JWT token",
			})
			ctx.Abort()
			return
		}
		claims:=token.Claims.(jwt.MapClaims)
		userId,ok:=claims["user_id"].(string)
		if !ok{
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid user id",
			})
			ctx.Abort()
			return
		}
		ctx.Set("user_id",userId)

		ctx.Next()

	}
}
