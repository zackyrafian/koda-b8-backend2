package middleware

import (
	"koda-b8-backend1/internal/libs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	  tokenString := c.GetHeader("Authorization")
		if tokenString ==  ""{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):]
		err := libs.VerifyToken(tokenString)
		if err != nil { 
		  c.JSON(http.StatusUnauthorized, gin.H{ 
				"error": "Unauthorized",
			})
				return
		}
		c.Next()
	}
}