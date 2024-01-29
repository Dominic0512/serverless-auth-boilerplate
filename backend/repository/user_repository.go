package repository

import (
	"context"
	"fmt"

	"github.com/Dominic0512/serverless-auth-boilerplate/ent/user"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
	"github.com/Dominic0512/serverless-auth-boilerplate/model"
	"github.com/google/uuid"
)

type UserRepository struct {
	User *model.UserClient
}

func (ur UserRepository) Find(name string) ([]*model.UserEntity, error) {
	u, err := ur.User.Query().Where(user.NameContains(name)).All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to query all user: %w", err)
	}
	return u, nil
}

func (ur UserRepository) FindOne(id uuid.UUID) (*model.UserEntity, error) {
	u, err := ur.User.Query().Where(user.ID(id)).Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to query single user: %w", err)
	}

	return u, nil
}

func (ur UserRepository) Create(user model.UserEntity) (*model.UserEntity, error) {
	u, err := ur.User.Create().
		SetEmail(user.Email).
		SetPassword(*user.Password).
		SetPasswordSalt(*user.PasswordSalt).
		SetName(user.Name).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return u, nil
}

func (ur UserRepository) Update(id uuid.UUID, properties model.UserEntity) (*model.UserEntity, error) {
	u, err := ur.User.UpdateOneID(id).
		SetName(properties.Name).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return u, nil
}

func (ur UserRepository) Delete(id uuid.UUID) (bool, error) {
	err := ur.User.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return false, fmt.Errorf("failed to delete user: %w", err)
	}

	return true, nil
}

func NewUserRepository(db *database.Database) UserRepository {
	return UserRepository{User: db.Client.User}
}
