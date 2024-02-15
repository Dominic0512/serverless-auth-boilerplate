package controller

import (
	"fmt"
	"net/http"

	"github.com/Dominic0512/serverless-auth-boilerplate/controller/request"
	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
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
	users, err := uc.us.Find()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "List user failed.",
		})
	}

	c.JSON(200, gin.H{
		"message": "List user successfully",
		"users":   users,
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

	input := domain.CreateUserWithoutPasswordInput{
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
	uri := request.ManipulateUri{}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uc.v.Validate.Struct(uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := uc.us.FindByID(uri.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Find user failed.",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get user by id successfully",
		"user":    user,
	})
}

func (uc UserController) Update(c *gin.Context) {
	uri := request.ManipulateUri{}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	request := request.UpdateUserRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	input := domain.UpdateUserInput{
		ID:   uri.ID,
		Name: request.Name,
	}

	user, err := uc.us.Update(input)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Update user failed.",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Update user successfully",
		"user":    user,
	})
}

func (uc UserController) PartialUpdate(c *gin.Context) {
	uri := request.ManipulateUri{}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	request := request.UpdateUserRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	input := domain.UpdateUserInput{
		ID:   uri.ID,
		Name: request.Name,
	}

	user, err := uc.us.Update(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Partial update user failed.",
		})
	}

	c.JSON(200, gin.H{
		"message": "Partial update user successfully",
		"user":    user,
	})
}

func (uc UserController) Delete(c *gin.Context) {
	uri := request.ManipulateUri{}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	input := domain.MaunipulateUserInput{
		ID: uri.ID,
	}

	err := uc.us.Delete(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Partial update user failed.",
		})
	}

	c.JSON(200, gin.H{
		"message": "Delete user successfully",
	})
}
