package model

import "github.com/Dominic0512/serverless-auth-boilerplate/ent"

type UserEntity = ent.User

type User struct {
	Email string
}

type CreateUserInput struct {
	Email    string
	Password string
}

type UserRepositoryInterface interface {
	GetAllUsers(name string) ([]*UserEntity, error)
	GetUser() (*UserEntity, error)
	CreateUser(user UserEntity) (*UserEntity, error)
	UpdateUser() (*UserEntity, error)
	DeleteUser() (*UserEntity, error)
}
