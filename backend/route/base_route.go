package route

import (
	"fmt"

	"github.com/Dominic0512/serverless-auth-boilerplate/route/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type BaseRoute struct {
	router                  *gin.Engine
	authMiddleware          middleware.AuthMiddleware
	errorHandlingMiddleware middleware.ErrorHandlingMiddleware
}

func (br BaseRoute) RegisterBeforeMiddlewares(fns ...gin.HandlerFunc) {
	for i := range fns {
		fn := fns[i]
		br.router.Use(func(c *gin.Context) {
			fn(c)
			c.Next()
		})
	}
}

func (br BaseRoute) RegisterAfterMiddlewares(fns ...gin.HandlerFunc) {
	// NOTE: Because gin middleware after next function is LIFO, we need to reverse the order of the fns
	for i := range fns {
		fn := fns[len(fns)-1-i]
		br.router.Use(func(c *gin.Context) {
			c.Next()
			fn(c)
		})
	}
}

func (br BaseRoute) Setup() {
	br.router.GET("/health", func(c *gin.Context) {
		fmt.Println("in health")
		c.JSON(200, gin.H{
			"message": "Alive...",
		})
	})

	br.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// br.RegisterBeforeMiddlewares(br.authMiddleware.OAuthTokenGuard)
	br.RegisterAfterMiddlewares(br.errorHandlingMiddleware.ErrorHandler)
}

func NewBaseRoute(
	router *gin.Engine,
	authMiddleware middleware.AuthMiddleware,
	errorHandlingMiddleware middleware.ErrorHandlingMiddleware,
) BaseRoute {
	return BaseRoute{
		router:                  router,
		authMiddleware:          authMiddleware,
		errorHandlingMiddleware: errorHandlingMiddleware,
	}
}
