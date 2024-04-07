package service

import (
	"context"
	"log"
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

func (as AuthService) GenerateAuthURL() string {
	url, err := as.auth.GenerateAuthCodeURL()
	if err != nil {
		log.Fatalf("Can not generate auth code url: %v", err)
		return ""
	}

	return url
}

func (as AuthService) SignUp(input domain.OAuthSignUpInput) (*string, error) {
	data, err := as.auth.ExchangeMetaDataByCode(input.Code)
	if err != nil {
		log.Fatalf("Can not exchange meta data by code: %v", err)
		return nil, err
	}

	ctx := context.Background()
	err = as.txHelper.WithTx(ctx, func(tx database.Tx) error {
		userProps := domain.UserEntity{
			Name:  strings.Split(data.Email, "@")[0],
			Email: data.Email,
		}
		u, err := as.userRepo.Create(ctx, tx, userProps)
		if err != nil {
			log.Fatalf("Can not create user properly: %v", err)
			return err
		}

		userProviderProps := domain.UserProviderEntity{
			Name:   domain.UserProviderNamePrimary,
			UserID: u.ID,
		}
		var _ domain.UserProviderEntity
		_, err = as.userProviderRepo.Create(ctx, tx, userProviderProps)
		if err != nil {
			log.Fatalf("Can not create user provider properly: %v", err)
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &data.AccessToken, nil
}

func (as AuthService) Login() {
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
