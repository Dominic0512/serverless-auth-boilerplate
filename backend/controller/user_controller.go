package controller

import (
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
	c.JSON(200, gin.H{
		"message": "Create user successfully",
	})
}

func (uc UserController) GetById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get user by id successfully",
	})
}

func (uc UserController) Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update user successfully",
	})
}

func (uc UserController) Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete user successfully",
	})
}
