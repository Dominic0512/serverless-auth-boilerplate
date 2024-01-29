package controller

import (
	"net/http"

	"github.com/Dominic0512/serverless-auth-boilerplate/controller/request"
	"github.com/Dominic0512/serverless-auth-boilerplate/model"
	"github.com/Dominic0512/serverless-auth-boilerplate/pkg/validate"
	"github.com/Dominic0512/serverless-auth-boilerplate/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	us service.UserService
	v  *validate.Validator
}

func NewUserController(us service.UserService, v *validate.Validator) UserController {
	return UserController{us, v}
}

func (uc UserController) List(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List user successfully",
	})
}

func (uc UserController) Create(c *gin.Context) {
	request := request.CreateUserRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uc.v.Validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	input := model.CreateUserWithoutPasswordInput{
		Email: request.Email,
	}

	user, err := uc.us.CreateWithoutPassword(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Create user failed.",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Create user successfully",
		"user":    user,
	})
}

func (uc UserController) GetById(c *gin.Context) {
	request := request.ManipulateRequest{}

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uc.v.Validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := uc.us.FindByID(request.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Create user failed.",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get user by id successfully",
		"user":    user,
	})
}

func (uc UserController) Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update user successfully",
	})
}

func (uc UserController) PartialUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Partial-update user successfully",
	})
}

func (uc UserController) Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete user successfully",
	})
}
