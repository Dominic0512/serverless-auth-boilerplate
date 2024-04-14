package repository

import (
	"context"
	"fmt"

	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
)

type UserProviderRepository struct {
	UserProvider *domain.UserProviderClient
}

func (upr UserProviderRepository) Create(ctx context.Context, tx database.Tx, properties domain.UserProviderEntity) (*domain.UserProviderEntity, error) {
	repo := upr.UserProvider

	if tx != nil {
		repo = tx.UserProvider
	}

	mutate := repo.Create().
		SetName(properties.Name).
		SetUserID(properties.UserID)

	userProvider, err := mutate.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user provider: %w", err)
	}

	return userProvider, nil
}

func NewUserProviderRepository(db *database.PSQLDatabase) *UserProviderRepository {
	return &UserProviderRepository{
		UserProvider: db.Client.UserProvider,
	}
}
