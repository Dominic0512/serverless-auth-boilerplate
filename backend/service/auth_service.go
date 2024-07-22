package service

import (
	"context"
	"strings"

	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/authenticator"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
	"github.com/Dominic0512/serverless-auth-boilerplate/pkg/helper"
)

type AuthService struct {
	txHelper         helper.TxHelper
	userRepo         domain.UserRepository
	userProviderRepo domain.UserProviderRepository
	auth             authenticator.Authenticator
}

func (as AuthService) exchangeMetaDataByCode(code string) (*authenticator.AuthMetaData, error) {
	data, err := as.auth.ExchangeMetaDataByCode(code)
	if err != nil {
		return nil, &domain.AuthorizationError{
			Message: err.Error(),
		}
	}

	if !data.EmailVerified {
		return nil, &domain.InvalidError{
			Entity:  "Email",
			Message: "Email is not verified",
		}
	}

	return data, nil
}

func (as AuthService) doUserCreationWithProvider(ctx context.Context, tx database.Tx, data *authenticator.AuthMetaData) error {
	userProps := domain.UserEntity{
		Name:  strings.Split(data.Email, "@")[0],
		Email: data.Email,
	}
	user, err := as.userRepo.Create(ctx, tx, userProps)
	if err != nil {
		return &domain.CreationError{
			Entity:  "User",
			Message: err.Error(),
		}
	}

	userProviderProps := domain.UserProviderEntity{
		Name:   domain.UserProviderName(data.Provider),
		UserID: user.ID,
	}

	var _ domain.UserProviderEntity
	_, err = as.userProviderRepo.Create(ctx, tx, userProviderProps)
	if err != nil {
		return &domain.CreationError{
			Entity:  "UserProvider",
			Message: err.Error(),
		}
	}

	return nil
}

func (as AuthService) GenerateAuthURL() (*string, error) {
	url, err := as.auth.GenerateAuthCodeURL()
	if err != nil {
		return nil, &domain.CreationError{
			Entity:  "AuthCodeURL",
			Message: err.Error(),
		}
	}

	return &url, nil
}

func (as AuthService) SignUp(input domain.OAuthSignUpInput) (*string, error) {
	data, err := as.exchangeMetaDataByCode(input.Code)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = as.txHelper.WithTx(ctx, func(tx database.Tx) error {
		return as.doUserCreationWithProvider(ctx, tx, data)
	})
	if err != nil {
		return nil, err
	}

	return &data.AccessToken, nil
}

func (as AuthService) SignIn(input domain.OAuthSignInInput) (*string, error) {
	data, err := as.exchangeMetaDataByCode(input.Code)
	if err != nil {
		return nil, err
	}

	_, err = as.userRepo.FindOneByEmail(context.Background(), data.Email)
	if err != nil {
		return nil, &domain.NotFoundError{
			Entity: "User",
			ID:     data.Email,
		}
	}

	return &data.AccessToken, nil
}

func NewAuthService(
	txHelper helper.TxHelper,
	userRepo domain.UserRepository,
	userProviderRepo domain.UserProviderRepository,
	auth authenticator.Authenticator,
) AuthService {
	return AuthService{
		txHelper,
		userRepo,
		userProviderRepo,
		auth,
	}
}
