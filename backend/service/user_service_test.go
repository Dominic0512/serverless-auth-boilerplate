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
	repo := new(mocks.UserRepository)
	pwh := new(mocks.PasswordHelper)

	repo.On("Find").Return(func() (users []*domain.UserEntity) {
		return []*domain.UserEntity{}
	}, func() error {
		return nil
	})

	userService := NewUserService(repo, pwh)

	t.Run("it should return result even there are nothing", func(t *testing.T) {
		users, err := userService.Find()
		if err != nil {
			t.Fatalf("user service find users error: %v", err)
		}

		assert.Equal(t, 0, len(users))
	})
}

func TestUserService_FindOne(t *testing.T) {
	dummyID := uuid.New()
	repo := new(mocks.UserRepository)
	pwh := new(mocks.PasswordHelper)

	repo.On("FindOne", mock.AnythingOfType("UUID")).Return(
		func(id uuid.UUID) *domain.UserEntity {
			return &domain.UserEntity{ID: dummyID}
		},
		func(id uuid.UUID) error {
			return nil
		},
	)

	userService := NewUserService(repo, pwh)

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
	repo := new(mocks.UserRepository)
	pwh := new(mocks.PasswordHelper)

	pwh.On("Hash", mock.AnythingOfType("string")).Return(
		func(password string) string {
			return ""
		},
		func(password string) error {
			return nil
		},
	)

	repo.On("Create", mock.AnythingOfType("User")).Return(
		func(user domain.UserEntity) *domain.UserEntity {
			return &domain.UserEntity{
				Email:    user.Email,
				Password: user.Password,
			}
		},
		func(user domain.UserEntity) error {
			return nil
		},
	)

	userService := NewUserService(repo, pwh)

	t.Run("it should create a new user with password", func(t *testing.T) {
		input := domain.CreateUserInput{
			Email:    "dummy@gmail.com",
			Password: "1",
		}
		user, err := userService.Create(input)

		a := assert.New(t)
		a.NotNil(user.Password)
		a.Nil(err)
	})
}

func TestUserService_CreateWithoutPassword(t *testing.T) {
	repo := new(mocks.UserRepository)
	pwh := new(mocks.PasswordHelper)

	repo.On("Create", mock.AnythingOfType("User")).Return(
		func(user domain.UserEntity) *domain.UserEntity {
			return &domain.UserEntity{
				Email: user.Email,
			}
		},
		func(user domain.UserEntity) error {
			return nil
		},
	)

	userService := NewUserService(repo, pwh)

	t.Run("it should create a new user without password", func(t *testing.T) {
		input := domain.CreateUserWithoutPasswordInput{
			Email: "dummy@gmail.com",
		}
		user, err := userService.CreateWithoutPassword(input)

		a := assert.New(t)
		a.Nil(user.Password)
		a.Nil(err)
	})
}

func TestUserService_Update(t *testing.T) {
	dummyID := uuid.New()
	repo := new(mocks.UserRepository)
	pwh := new(mocks.PasswordHelper)

	repo.On("Update", mock.AnythingOfType("UUID"), mock.AnythingOfType("User")).Return(
		func(id uuid.UUID, user domain.UserEntity) *domain.UserEntity {
			return &domain.UserEntity{ID: dummyID, Name: user.Name}
		},
		func(id uuid.UUID, user domain.UserEntity) error {
			return nil
		},
	)

	userService := NewUserService(repo, pwh)

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
	repo := new(mocks.UserRepository)
	pwh := new(mocks.PasswordHelper)

	repo.On("Delete", mock.AnythingOfType("UUID")).Return(
		func(id uuid.UUID) error {
			return nil
		},
	)

	userService := NewUserService(repo, pwh)

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
