package controller

import (
	"net/http"

	"github.com/Dominic0512/serverless-auth-boilerplate/controller/request"
	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/Dominic0512/serverless-auth-boilerplate/pkg/validate"
	"github.com/Dominic0512/serverless-auth-boilerplate/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	us service.UserService
	v  *validate.Validator
}

func NewAuthController(us service.UserService, v *validate.Validator) AuthController {
	return AuthController{us, v}
}

func (ac AuthController) SignIn(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Sign-in successfully",
	})
}

func (ac AuthController) SignUp(c *gin.Context) {
	request := request.SignUpRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ac.v.Validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := domain.CreateUserInput{
		Email:    request.Email,
		Password: request.Password,
	}

	u, err := ac.us.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Create user failed.",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Sign-up successfully",
		"user":    u,
	})
}
