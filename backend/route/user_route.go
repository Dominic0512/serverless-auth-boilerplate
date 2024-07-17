package route

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/controller"
	"github.com/Dominic0512/serverless-auth-boilerplate/route/middleware"
	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	router         *gin.Engine
	authMiddleware middleware.AuthMiddleware
	userController controller.UserController
}

func (ur UserRoute) Setup() {
	user := ur.router.Group("/api/users")
	user.Use(ur.authMiddleware.OAuthTokenGuard)
	{
		user.GET("", ur.userController.List)
		user.POST("", ur.userController.Create)
		user.GET("/:id", ur.userController.GetById)
		user.PUT("/:id", ur.userController.Update)
		user.PATCH("/:id", ur.userController.PartialUpdate)
		user.DELETE("/:id", ur.userController.Delete)
	}
}

func NewUserRoute(
	router *gin.Engine,
	authMiddleware middleware.AuthMiddleware,
	userController controller.UserController,
) UserRoute {
	return UserRoute{
		router:         router,
		authMiddleware: authMiddleware,
		userController: userController,
	}
}
