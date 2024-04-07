package service

import (
	"testing"

	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/Dominic0512/serverless-auth-boilerplate/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_Find(t *testing.T) {
	txHelper := new(mocks.TxHelper)
	userRepo := new(mocks.UserRepository)
	userProviderRepo := new(mocks.UserProviderRepository)
	pwh := new(mocks.PasswordHelper)

	userRepo.On("Find", mock.AnythingOfType("backgroundCtx")).Return([]*domain.UserEntity{}, nil)

	userService := NewUserService(txHelper, userRepo, userProviderRepo, pwh)

	t.Run("it should return result even there are nothing", func(t *testing.T) {
		users, err := userService.Find()

		a := assert.New(t)
		a.Nil(err)
		a.Equal(0, len(users))
	})
}

func TestUserService_FindOne(t *testing.T) {
	dummyID := uuid.New()
	txHelper := new(mocks.TxHelper)
	userRepo := new(mocks.UserRepository)
	userProviderRepo := new(mocks.UserProviderRepository)
	pwh := new(mocks.PasswordHelper)

	userRepo.On("FindOne", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("UUID")).Return(
		&domain.UserEntity{ID: dummyID},
		nil,
	)

	userService := NewUserService(txHelper, userRepo, userProviderRepo, pwh)

	t.Run("it should return the result", func(t *testing.T) {
		user, err := userService.FindByID(dummyID.String())
		a := assert.New(t)
		a.Equal(dummyID, user.ID)
		a.Nil(err)
	})

	t.Run("it should throw invalid id type error", func(t *testing.T) {
		user, err := userService.FindByID("random")

		a := assert.New(t)
		a.Nil(user)
		a.ErrorContains(err, "invalid id type:")
	})
}

func TestUserService_Create(t *testing.T) {
	txHelper := new(mocks.TxHelper)
	userRepo := new(mocks.UserRepository)
	userProviderRepo := new(mocks.UserProviderRepository)
	pwh := new(mocks.PasswordHelper)

	pwh.On("Hash", mock.AnythingOfType("string")).Return(
		func(password string) string {
			return ""
		},
		func() error {
			return nil
		},
	)

	userRepo.On("Create", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("Tx"), mock.AnythingOfType("User")).Return(
		func(user domain.UserEntity) *domain.UserEntity {
			return &domain.UserEntity{
				Email:    user.Email,
				Password: user.Password,
			}
		},
		func() error {
			return nil
		},
	).On("FindOneByEmail", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("string")).Return(
		&domain.UserEntity{
			Email: "dummy@gmail.com",
		},
		nil,
	)

	userProviderRepo.On("Create", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("Tx"), mock.AnythingOfType("UserProvider")).Return(
		func(userProvider domain.UserProviderEntity) *domain.UserProviderEntity {
			return &domain.UserProviderEntity{
				Picture: "dummy picture",
				Name:    domain.UserProviderNamePrimary,
			}
		},
		func() error {
			return nil
		},
	)

	txHelper.On("WithTx", mock.Anything, mock.Anything).Return(nil)

	userService := NewUserService(txHelper, userRepo, userProviderRepo, pwh)

	t.Run("it should create a new user with password", func(t *testing.T) {
		input := domain.CreateUserInput{
			Email:    "dummy@gmail.com",
			Password: "1",
		}
		user, err := userService.Create(input)

		a := assert.New(t)
		a.NotNil(user)
		a.Nil(err)
	})
}

func TestUserService_CreateWithoutPassword(t *testing.T) {
	txHelper := new(mocks.TxHelper)
	userRepo := new(mocks.UserRepository)
	userProviderRepo := new(mocks.UserProviderRepository)
	pwh := new(mocks.PasswordHelper)

	userRepo.On("Create", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("Tx"), mock.AnythingOfType("User")).Return(
		func(user domain.UserEntity) *domain.UserEntity {
			return &domain.UserEntity{
				Email: user.Email,
			}
		},
		func(user domain.UserEntity) error {
			return nil
		},
	).On("FindOneByEmail", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("string")).Return(
		&domain.UserEntity{
			Email: "dummy@gmail.com",
		},
		nil,
	)

	userProviderRepo.On("Create", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("Tx"), mock.AnythingOfType("UserProvider")).Return(
		func(userProvider domain.UserProviderEntity) *domain.UserProviderEntity {
			return &domain.UserProviderEntity{
				Picture: "dummy picture",
				Name:    domain.UserProviderNamePrimary,
			}
		},
		func() error {
			return nil
		},
	)

	txHelper.On("WithTx", mock.Anything, mock.Anything).Return(nil)

	userService := NewUserService(txHelper, userRepo, userProviderRepo, pwh)

	t.Run("it should create a new user without password", func(t *testing.T) {
		input := domain.CreateUserWithoutPasswordInput{
			Email: "dummy@gmail.com",
		}
		user, err := userService.CreateWithoutPassword(input)

		a := assert.New(t)
		a.NotNil(user)
		a.Nil(err)
	})
}

func TestUserService_Update(t *testing.T) {
	dummyID := uuid.New()
	txHelper := new(mocks.TxHelper)
	userRepo := new(mocks.UserRepository)
	userProviderRepo := new(mocks.UserProviderRepository)
	pwh := new(mocks.PasswordHelper)

	userRepo.On("Update", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("UUID"), mock.AnythingOfType("User")).Return(
		&domain.UserEntity{
			ID:   dummyID,
			Name: "modified",
		},
		nil,
	)

	userService := NewUserService(txHelper, userRepo, userProviderRepo, pwh)

	t.Run("it should return the result with updatable fields", func(t *testing.T) {
		input := domain.UpdateUserInput{
			ID:   dummyID.String(),
			Name: "modified",
		}
		user, err := userService.Update(input)
		a := assert.New(t)
		a.Equal(dummyID, user.ID)
		a.Equal(input.Name, user.Name)
		a.Nil(err)
	})

	t.Run("it should throw invalid id type error", func(t *testing.T) {
		input := domain.UpdateUserInput{
			ID:   "random",
			Name: "modified",
		}
		user, err := userService.Update(input)

		a := assert.New(t)
		a.Nil(user)
		a.ErrorContains(err, "invalid id type:")
	})
}

func TestUserService_Delete(t *testing.T) {
	dummyID := uuid.New()
	txHelper := new(mocks.TxHelper)
	userRepo := new(mocks.UserRepository)
	userProviderRepo := new(mocks.UserProviderRepository)
	pwh := new(mocks.PasswordHelper)

	userRepo.On("Delete", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("UUID")).Return(nil)

	userService := NewUserService(txHelper, userRepo, userProviderRepo, pwh)

	t.Run("it should return the result", func(t *testing.T) {
		input := domain.MaunipulateUserInput{
			ID: dummyID.String(),
		}
		err := userService.Delete(input)

		a := assert.New(t)
		a.Nil(err)
	})

	t.Run("it should throw invalid id type error", func(t *testing.T) {
		input := domain.MaunipulateUserInput{
			ID: "random",
		}
		err := userService.Delete(input)

		a := assert.New(t)
		a.ErrorContains(err, "invalid id type:")
	})
}
