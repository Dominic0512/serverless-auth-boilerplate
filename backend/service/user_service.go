package service

import (
	"fmt"
	"strings"

	"github.com/Dominic0512/serverless-auth-boilerplate/model"
	"github.com/Dominic0512/serverless-auth-boilerplate/repository"
	"github.com/google/uuid"
)

type UserService struct {
	repo repository.UserRepository
}

func (us UserService) Find(name string) ([]*model.UserEntity, error) {
	users, err := us.repo.Find(name)
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}
	return users, nil
}

func (us UserService) FindByID(id string) (*model.UserEntity, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id type: %w", err)
	}

	user, err := us.repo.FindOne(uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}
	return user, nil
}

func (us UserService) Create(input model.CreateUserInput) (*model.UserEntity, error) {
	salt := "test"
	userProps := model.UserEntity{
		Name:         strings.Split(input.Email, "@")[0],
		Email:        input.Email,
		Password:     &input.Password,
		PasswordSalt: &salt,
	}

	user, err := us.repo.Create(userProps)
	if err != nil {
		return nil, fmt.Errorf("failed mutating user: %w", err)
	}

	return user, nil
}

func (us UserService) CreateWithoutPassword(input model.CreateUserWithoutPasswordInput) (*model.UserEntity, error) {
	userProps := model.UserEntity{
		Name:  strings.Split(input.Email, "@")[0],
		Email: input.Email,
	}

	user, err := us.repo.Create(userProps)
	if err != nil {
		return nil, fmt.Errorf("failed create user without password: %w", err)
	}

	return user, nil
}

func (us UserService) Update(input model.UpdateUserInput) (*model.UserEntity, error) {
	user, err := us.repo.Update(input.ID, model.UserEntity{
		Name: input.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed create user without password: %w", err)
	}
	return user, nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo}
}
