package controller

import (
	"net/http"

	"github.com/Dominic0512/serverless-auth-boilerplate/controller/request"
	"github.com/Dominic0512/serverless-auth-boilerplate/controller/response"
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

// GenerateAuthURL godoc
// @Summary Generate oauth login url
// @Schemes http
// @Description Currently, the authorization is integrated with Auth0. This endpoint will generate an authorization URL for the client to redirect to the Auth0 login page.
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} response.GenerateAuthURLResponse "ok"
// @Router /oauth-url [get]
func (ac AuthController) GenerateAuthURL(c *gin.Context) {
	url, err := ac.as.GenerateAuthURL()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Generate auth url failed.",
		})
		return
	}

	c.JSON(200, response.GenerateAuthURLResponse{
		Url: url,
	})
}

// SignIn godoc
// @Summary SignIn with oauth code
// @Schemes http
// @Description SignIn with oauth code
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} response.TokenResponse "ok"
// @Router /sign-in [post]
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

	c.JSON(http.StatusAccepted, response.TokenResponse{
		Token:      token.AccessToken,
		Token_type: token.TokenType,
	})
}

// SignUp godoc
// @Summary SignUp with oauth code
// @Schemes http
// @Description SignUp with oauth code
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} response.TokenResponse "ok"
// @Router /sign-up [post]
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

	c.JSON(http.StatusAccepted, response.TokenResponse{
		Token:      token.AccessToken,
		Token_type: token.TokenType,
	})
}
