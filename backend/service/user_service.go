package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
	"github.com/Dominic0512/serverless-auth-boilerplate/pkg/helper"
	"github.com/google/uuid"
)

type UserService struct {
	txHelper         helper.TxHelper
	userRepo         domain.UserRepository
	userProviderRepo domain.UserProviderRepository
	pwh              helper.PasswordHelper
}

func (us UserService) Find() ([]*domain.UserEntity, error) {
	users, err := us.userRepo.Find(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}
	return users, nil
}

func (us UserService) FindByID(id string) (*domain.UserEntity, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id type: %w", err)
	}

	user, err := us.userRepo.FindOne(context.Background(), uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}
	return user, nil
}

func (us UserService) Create(input domain.CreateUserInput) (*domain.UserEntity, error) {
	ctx := context.Background()

	err := us.txHelper.WithTx(ctx, func(tx database.Tx) error {
		hashedPassword, err := us.pwh.Hash(input.Password)
		if err != nil {
			return fmt.Errorf("failed to hash the password: %w", err)
		}

		userProps := domain.UserEntity{
			Name:     strings.Split(input.Email, "@")[0],
			Email:    input.Email,
			Password: &hashedPassword,
		}

		user, err := us.userRepo.Create(ctx, tx, userProps)
		if err != nil {
			return fmt.Errorf("failed mutating user: %w", err)
		}

		userProviderProps := domain.UserProviderEntity{
			Name:   domain.UserProviderNamePrimary,
			UserID: user.ID,
		}

		var _ domain.UserProviderEntity
		_, err = us.userProviderRepo.Create(ctx, tx, userProviderProps)
		if err != nil {
			return fmt.Errorf("can not create user provider properly: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	user, err := us.userRepo.FindOneByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us UserService) CreateWithoutPassword(input domain.CreateUserWithoutPasswordInput) (*domain.UserEntity, error) {
	ctx := context.Background()

	err := us.txHelper.WithTx(ctx, func(tx database.Tx) error {
		userProps := domain.UserEntity{
			Name:  strings.Split(input.Email, "@")[0],
			Email: input.Email,
		}

		user, err := us.userRepo.Create(ctx, tx, userProps)
		if err != nil {
			return fmt.Errorf("failed create user without password: %v", err)
		}

		userProviderProps := domain.UserProviderEntity{
			Name:   domain.UserProviderNamePrimary,
			UserID: user.ID,
		}

		var _ domain.UserProviderEntity
		_, err = us.userProviderRepo.Create(ctx, tx, userProviderProps)
		if err != nil {
			return fmt.Errorf("can not create user provider properly: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	user, err := us.userRepo.FindOneByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us UserService) Update(input domain.UpdateUserInput) (*domain.UserEntity, error) {
	uuid, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid id type: %v", err)
	}

	user, err := us.userRepo.Update(context.Background(), uuid, domain.UserEntity{
		Name: input.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update by id: %v", err)
	}

	return user, nil
}

func (us UserService) Delete(input domain.MaunipulateUserInput) error {
	uuid, err := uuid.Parse(input.ID)
	if err != nil {
		return fmt.Errorf("invalid id type: %v", err)
	}

	err = us.userRepo.Delete(context.Background(), uuid)
	if err != nil {
		return fmt.Errorf("failed to delete user by id: %v", err)
	}

	return nil
}

func NewUserService(
	txHelper helper.TxHelper,
	userRepo domain.UserRepository,
	userProviderRepo domain.UserProviderRepository,
	pwh helper.PasswordHelper,
) UserService {
	return UserService{
		txHelper,
		userRepo,
		userProviderRepo,
		pwh,
	}
}
