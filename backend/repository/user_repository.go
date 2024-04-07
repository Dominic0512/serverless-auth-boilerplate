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

func (ur UserRepository) Find(ctx context.Context) ([]*domain.UserEntity, error) {
	u, err := ur.User.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query all user: %w", err)
	}
	return u, nil
}

func (ur UserRepository) FindOne(ctx context.Context, id uuid.UUID) (*domain.UserEntity, error) {
	u, err := ur.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query single user: %w", err)
	}

	return u, nil
}

func (ur UserRepository) FineOneByEmail(ctx context.Context, email string) (*domain.UserEntity, error) {
	u, err := ur.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query single user: %w", err)
	}

	return u, nil
}

func (ur UserRepository) Create(ctx context.Context, tx database.Tx, user domain.UserEntity) (*domain.UserEntity, error) {
	repo := ur.User

	if tx != nil {
		repo = tx.User
	}

	mutate := repo.Create().
		SetEmail(user.Email).
		SetName(user.Name)

	if user.Password != nil {
		mutate.SetPassword(*user.Password)
	}

	u, err := mutate.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return u, nil
}

func (ur UserRepository) Update(ctx context.Context, id uuid.UUID, properties domain.UserEntity) (*domain.UserEntity, error) {
	u, err := ur.User.UpdateOneID(id).
		SetName(properties.Name).
		Save(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return u, nil
}

func (ur UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := ur.User.DeleteOneID(id).Exec(ctx)
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
