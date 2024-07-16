package middleware

import (
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct{}

func (am AuthMiddleware) OAuthTokenGuard(c *gin.Context) {
	if len(c.Errors) > 0 {
		c.JSON(401, gin.H{
			"message": c.Errors.String(),
		})
		c.Abort()
	}
}

func NewAuthMiddleware() AuthMiddleware {
	return AuthMiddleware{}
}
