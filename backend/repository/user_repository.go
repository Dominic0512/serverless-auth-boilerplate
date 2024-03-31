package repository

import (
	"context"
	"fmt"

	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/Dominic0512/serverless-auth-boilerplate/ent/user"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
	"github.com/google/uuid"
)

type UserRepository struct {
	User *domain.UserClient
}

func (ur UserRepository) Find() ([]*domain.UserEntity, error) {
	u, err := ur.User.Query().All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to query all user: %w", err)
	}
	return u, nil
}

func (ur UserRepository) FindOne(id uuid.UUID) (*domain.UserEntity, error) {
	u, err := ur.User.Query().Where(user.ID(id)).Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to query single user: %w", err)
	}

	return u, nil
}

func (ur UserRepository) Create(user domain.UserEntity) (*domain.UserEntity, error) {
	mutate := ur.User.Create().
		SetEmail(user.Email).
		SetName(user.Name)

	if user.Password != nil {
		mutate.SetPassword(*user.Password)
	}

	u, err := mutate.Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return u, nil
}

func (ur UserRepository) Update(id uuid.UUID, properties domain.UserEntity) (*domain.UserEntity, error) {
	u, err := ur.User.UpdateOneID(id).
		SetName(properties.Name).
		Save(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return u, nil
}

func (ur UserRepository) Delete(id uuid.UUID) error {
	err := ur.User.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func NewUserRepository(db *database.Database) *UserRepository {
	return &UserRepository{
		User: db.Client.User,
	}
}
