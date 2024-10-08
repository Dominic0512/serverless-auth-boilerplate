package route

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/controller/auth"
	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	router         *gin.Engine
	authController auth.AuthController
}

func (ar AuthRoute) Setup() {
	auth := ar.router.Group("/api/v1/auth")
	{
		auth.GET("/oauth-url", ar.authController.GenerateAuthURL)
		auth.POST("/sign-in", ar.authController.SignIn)
		auth.POST("/sign-up", ar.authController.SignUp)
	}
}

func NewAuthRoute(
	router *gin.Engine,
	authController auth.AuthController,
) AuthRoute {
	return AuthRoute{
		router:         router,
		authController: authController,
	}
}
