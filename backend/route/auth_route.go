package route

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/controller"
	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	router         *gin.Engine
	authController controller.AuthController
}

func (ar AuthRoute) Setup() {
	auth := ar.router.Group("/api")
	{
		auth.POST("/sign-in", ar.authController.SignIn)
		auth.POST("/sign-up", ar.authController.SignUp)
	}
}

func NewAuthRoute(
	router *gin.Engine,
	authController controller.AuthController,
) AuthRoute {
	return AuthRoute{
		router:         router,
		authController: authController,
	}
}
