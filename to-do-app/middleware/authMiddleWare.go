package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"to-do-app/models"
	"net/http"
	"to-do-app/utils"
)

func AuthMiddlewareToken(roles ...models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			log.Println("Token not found")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			log.Println("Unauthorized access")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		for _, role := range roles {
			if role == claims.UserRole {
				log.Println("Authorized access")
				c.Set("userId", claims.UserId)
				c.Set("role", claims.UserRole)
				c.Next()
				return
			}
		}

		log.Println("Forbidden access", roles, claims.UserRole)
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden access"})
		c.Abort()
		return
	}
}

func authMiddleWarePassword() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			log.Println("Token not found")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			log.Println("Unauthorized access")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserId)
		c.Set("role", claims.UserRole)
		c.Next()
	}
}