package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement your authentication logic here
		authheader := c.GetHeader("Authorization")
		if authheader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}
		// Validate the token and set user information in the context if needed

		tostring := strings.TrimPrefix(authheader, "Bearer ")
		if tostring == authheader {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		_, err := jwt.ParseWithClaims(tostring, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method and return the secret key
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte("your_secret_key"), nil
		})
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}

}
