package middleware

import (
	"fmt"
	"ginapp/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Request Headers", c.Request.Header) //helps in debugging request is being made

		authheader := c.GetHeader("Authorization") // contains authentication credentials
		tokenString := helper.GetTokenFromHeader(authheader)

		if tokenString == "" {
			var err error
			tokenString, err = c.Cookie("Authorization")
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		userId, userEmail, err := helper.ExtractUserIDFromToken(tokenString)
		fmt.Println("userId", userId, "Useremail", userEmail)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user_id", userId)
		c.Set("user_email", userEmail)
		c.Next()
	}
}
