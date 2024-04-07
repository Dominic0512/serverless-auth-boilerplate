package domain

import (
	"context"

	"github.com/Dominic0512/serverless-auth-boilerplate/ent"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
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

type MaunipulateUserInput struct {
	ID string
}

type UpdateUserInput struct {
	ID   string
	Name string
}

type UserRepository interface {
	Find(ctx context.Context) ([]*UserEntity, error)
	FindOne(ctx context.Context, id uuid.UUID) (*UserEntity, error)
	FindOneByEmail(ctx context.Context, email string) (*UserEntity, error)
	Create(ctx context.Context, tx database.Tx, user UserEntity) (*UserEntity, error)
	Update(ctx context.Context, id uuid.UUID, properties UserEntity) (*UserEntity, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
