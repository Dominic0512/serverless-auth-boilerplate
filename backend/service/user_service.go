package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/Dominic0512/serverless-auth-boilerplate/ent"
	"github.com/Dominic0512/serverless-auth-boilerplate/model"
	"github.com/Dominic0512/serverless-auth-boilerplate/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (us UserService) GetAllUsers(name string) ([]*ent.User, error) {
	users, err := us.repo.GetAllUsers(name)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", users)
	return users, nil
}

func (us UserService) CreateUser(input model.CreateUserInput) (*ent.User, error) {
	salt := "test"
	userProps := ent.User{
		Name:         strings.Split(input.Email, "@")[0],
		Email:        input.Email,
		Password:     &input.Password,
		PasswordSalt: &salt,
	}

	user, err := us.repo.CreateUser(userProps)
	if err != nil {
		return nil, fmt.Errorf("failed mutating user: %w", err)
	}
	return user, nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo}
}
