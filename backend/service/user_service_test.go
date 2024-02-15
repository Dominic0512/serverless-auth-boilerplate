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

	repo.On("Find").Return(func() (users []*domain.UserEntity) {
		return []*domain.UserEntity{}
	}, func() error {
		return nil
	})

	userService := NewUserService(repo)

	t.Run("find user list", func(t *testing.T) {
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

	repo.On("FindOne", mock.AnythingOfType("UUID")).Return(func(id uuid.UUID) *domain.UserEntity {
		return &domain.UserEntity{ID: dummyID}
	}, func(id uuid.UUID) error {
		return nil
	})

	userService := NewUserService(repo)

	t.Run("find by formated uuid string", func(t *testing.T) {
		user, err := userService.FindByID(dummyID.String())
		a := assert.New(t)
		a.Equal(dummyID, user.ID)
		a.Nil(err)
	})

	t.Run("find by random string", func(t *testing.T) {
		user, err := userService.FindByID("random")

		a := assert.New(t)
		a.Nil(user)
		a.ErrorContains(err, "invalid id type:")
	})
}
