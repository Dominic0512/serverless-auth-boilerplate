package middleware

import (
	"errors"
	"net/http"

	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/gin-gonic/gin"
)

type ErrorHandlingMiddleware struct{}

func (ehm ErrorHandlingMiddleware) ErrorHandler(c *gin.Context) {
	lastErr := c.Errors.Last()
	if lastErr == nil {
		return
	}

	var domainErr domain.DomainError
	if errors.As(lastErr.Err, &domainErr) {
		switch e := lastErr.Err.(type) {
		case domain.InvalidError:
			c.JSON(http.StatusBadRequest, gin.H{"message": e.Error()})
		case domain.CreationError:
			c.JSON(http.StatusBadRequest, gin.H{"message": e.Error()})
		case domain.ValidationError:
			c.JSON(http.StatusBadRequest, gin.H{"message": e.Error()})
		case domain.AuthorizationError:
			c.JSON(http.StatusUnauthorized, gin.H{"message": e.Error()})
		case domain.NotFoundError:
			c.JSON(http.StatusNotFound, gin.H{"message": e.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service Unavailable Error"})
}

func (ehm ErrorHandlingMiddleware) ErrorRecoveryHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			c.Abort()
		}
	}()
}

func NewErrorHandlingMiddleware() ErrorHandlingMiddleware {
	return ErrorHandlingMiddleware{}
}
