package model

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/ent"
	"github.com/google/uuid"
)

type UserEntity = ent.User

type UserClient = ent.UserClient

type User struct {
	Email string
}

type CreateUserWithoutPasswordInput struct {
	Email string
}

type CreateUserInput struct {
	Email    string
	Password string
}

type UpdateUserInput struct {
	ID   uuid.UUID
	Name string
}

type UserRepositoryInterface interface {
	Find(name string) ([]*UserEntity, error)
	FindOne(id uuid.UUID) (*UserEntity, error)
	Create(user UserEntity) (*UserEntity, error)
	Update(id uuid.UUID, properties UserEntity) (*UserEntity, error)
	Delete(id uuid.UUID) (*UserEntity, error)
}
