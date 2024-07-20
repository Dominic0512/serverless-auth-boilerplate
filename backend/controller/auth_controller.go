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
	v  *validate.Validator
	as service.AuthService
	us service.UserService
}

func NewAuthController(
	v *validate.Validator,
	as service.AuthService,
	us service.UserService,
) AuthController {
	return AuthController{v, as, us}
}

func (ac AuthController) GenerateAuthURL(c *gin.Context) {
	url, err := ac.as.GenerateAuthURL()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Generate auth url failed.",
		})
		return
	}

	c.JSON(200, gin.H{
		"url": url,
	})
}

func (ac AuthController) SignIn(c *gin.Context) {
	request := request.SignInRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := ac.v.Validate.Struct(request); err != nil {
		c.Error(err)
		return
	}

	input := domain.OAuthSignInInput{
		Code: request.Code,
	}
	token, err := ac.as.SignIn(input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"token": token,
	})
}

func (ac AuthController) SignUp(c *gin.Context) {
	request := request.SignUpRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := ac.v.Validate.Struct(request); err != nil {
		c.Error(err)
		return
	}

	input := domain.OAuthSignUpInput{
		Code: request.Code,
	}
	token, err := ac.as.SignUp(input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"token": token,
	})
}
