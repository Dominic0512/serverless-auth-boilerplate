package user

import (
	"fmt"
	"net/http"

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

// List godoc
// @Id ListUsers
// @Summary List users
// @Schemes http
// @Description List users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} UsersResponse "ok"
// @Router /users [get]
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

// Create godoc
// @Id CreateUser
// @Summary Create user
// @Schemes http
// @Description Create user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} UserResponse "ok"
// @Router /users [post]
func (uc UserController) Create(c *gin.Context) {
	request := CreateUserRequest{}

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

	c.JSON(http.StatusAccepted, UserResponse{User: User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}})
}

// GetById godoc
// @Id GetUser
// @Summary Get user by id
// @Schemes http
// @Description Get user by id
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} UserResponse "ok"
// @Router /users/{id} [get]
func (uc UserController) GetById(c *gin.Context) {
	uri := ManipulateUri{}

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

	c.JSON(http.StatusAccepted, UserResponse{User: User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}})
}

// Update godoc
// @Id UpdateUser
// @Summary Update user
// @Schemes http
// @Description Update user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} UserResponse "ok"
// @Router /users/{id} [put]
func (uc UserController) Update(c *gin.Context) {
	uri := ManipulateUri{}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	request := UpdateUserRequest{}

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

	c.JSON(http.StatusAccepted, UserResponse{User: User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}})
}

// PartialUpdate godoc
// @Id PartialUpdateUser
// @Summary Partial update user
// @Schemes http
// @Description Partial update user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} UserResponse "ok"
// @Router /users/{id} [patch]
func (uc UserController) PartialUpdate(c *gin.Context) {
	uri := ManipulateUri{}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	request := UpdateUserRequest{}

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

	c.JSON(http.StatusAccepted, UserResponse{User: User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}})
}

// Delete godoc
// @Id DeleteUser
// @Summary Delete user
// @Schemes http
// @Description Delete user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Router /users/{id} [delete]
func (uc UserController) Delete(c *gin.Context) {
	uri := ManipulateUri{}

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
			"message": "Delete user failed.",
		})
	}

	c.JSON(200, gin.H{})
}
