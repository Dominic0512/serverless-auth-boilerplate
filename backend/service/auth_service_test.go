package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/authenticator"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
	"github.com/Dominic0512/serverless-auth-boilerplate/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_doUserCreationWithProvider(t *testing.T) {
	txHelper := new(mocks.TxHelper)
	userRepo := new(mocks.UserRepository)
	userProviderRepo := new(mocks.UserProviderRepository)
	auth := new(mocks.Authenticator)

	as := NewAuthService(txHelper, userRepo, userProviderRepo, auth)

	t.Run("it will success", func(t *testing.T) {
		userRepo.On("Create", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("Tx"), mock.AnythingOfType("User")).Return(
			&domain.UserEntity{
				Email: "dummy@gmail.com",
			}, nil,
		).Once()

		userProviderRepo.On("Create", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("Tx"), mock.AnythingOfType("UserProvider")).Return(
			&domain.UserProviderEntity{
				Picture: "dummy picture",
				Name:    domain.UserProviderNamePrimary,
			}, nil,
		).Once()

		var tx database.Tx
		err := as.doUserCreationWithProvider(context.Background(), tx, &authenticator.AuthMetaData{Email: "dummy@gmail.com"})
		a := assert.New(t)
		a.Nil(err)
	})

	t.Run("it will failure by user creation", func(t *testing.T) {
		userRepo.On("Create", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("Tx"), mock.AnythingOfType("User")).Return(
			nil, errors.New("user creation dummy error"),
		).Once()

		err := as.doUserCreationWithProvider(context.Background(), nil, &authenticator.AuthMetaData{Email: "dummy@gmail.com"})

		a := assert.New(t)
		a.Error(err)
	})

	t.Run("it will failure by user provider creation", func(t *testing.T) {
		userRepo.On("Create", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("Tx"), mock.AnythingOfType("User")).Return(
			&domain.UserEntity{
				ID:    uuid.New(),
				Email: "dummy@gmail.com",
			}, nil,
		).Once()

		userProviderRepo.On("Create", mock.AnythingOfType("backgroundCtx"), mock.AnythingOfType("Tx"), mock.AnythingOfType("UserProvider")).Return(
			nil, errors.New("user provider creation dummy error"),
		).Once()

		err := as.doUserCreationWithProvider(context.Background(), nil, &authenticator.AuthMetaData{Email: "dummy@gmail.com"})
		a := assert.New(t)
		a.Error(err)
	})
}

func TestAuthService_GenerateAuthURL(t *testing.T) {
	txHelper := new(mocks.TxHelper)
	userRepo := new(mocks.UserRepository)
	userProviderRepo := new(mocks.UserProviderRepository)
	auth := new(mocks.Authenticator)

	auth.On("GenerateAuthCodeURL").Return(
		func() string {
			return ""
		},
		func() error {
			return nil
		},
	)

	as := NewAuthService(txHelper, userRepo, userProviderRepo, auth)
	t.Run("it should return oauth url without err", func(t *testing.T) {
		_, err := as.auth.GenerateAuthCodeURL()

		a := assert.New(t)
		a.Nil(err)
	})
}

func TestAuthService_SignUp(t *testing.T) {
	dummyJWTToken := "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwiaXNzIjoiaHR0cHM6Ly9kZXYtd3A4b2h0dXFobXJvNW13YS51cy5hdXRoMC5jb20vIn0..fFxOTGtxT9EcM621.RuzzHt-ToBnPbD8mjnSqg6F9HglV2oeaQPToQPDcTGwUa_3DyPrM_2mCZ1sZxxndnqrsQG7RW_Nv1gBHuEYm0i2V8lTx2YS0GbI5YeYe_KS5STeHRqgQpPIGiIj3TgKLSIzE4UznBy_3vOIkFbnsj6bbJRVu9YV9OIiQUkxRlx38isZotjHC91m-XJj02-0ZOcgnBMLQQSLix0Ti-w332J15ViD85Ps3E6d26RqIiyaSzd2PR1kxb5ejmw2WC2VXgZBNBTmu3ZytiLunDYydK-HqxPrujLptU6-utqcCpo_UaPVNH2ahoSCDZiOyMdWnBtbFiCDEFRNKsmdBU6YU83tX.LLQgf8yZ6NiMHasatqanHw"
	txHelper := new(mocks.TxHelper)
	userRepo := new(mocks.UserRepository)
	userProviderRepo := new(mocks.UserProviderRepository)
	auth := new(mocks.Authenticator)

	txHelper.On("WithTx", mock.Anything, mock.Anything).Return(nil)

	auth.On("ExchangeMetaDataByCode", mock.AnythingOfType("string")).Return(
		func(code string) *authenticator.AuthMetaData {
			return &authenticator.AuthMetaData{
				AccessToken:   dummyJWTToken,
				Email:         "dummy@gmail.com",
				EmailVerified: true,
			}
		},
		func(code string) error {
			return nil
		},
	)

	as := NewAuthService(txHelper, userRepo, userProviderRepo, auth)

	t.Run("it will return access token", func(t *testing.T) {
		input := domain.OAuthSignUpInput{
			Code: "",
		}
		token, err := as.SignUp(input)

		a := assert.New(t)
		a.Equal(dummyJWTToken, *token)
		a.Nil(err)
	})
}
