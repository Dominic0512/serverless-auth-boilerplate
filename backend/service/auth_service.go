package service

import (
	"log"
	"strings"

	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/authenticator"
)

type AuthService struct {
	userRepo         domain.UserRepository
	userProviderRepo domain.UserProviderRepository
	auth             authenticator.Authenticator
}

func (as AuthService) GenerateAuthURL() string {
	url, err := as.auth.GenerateAuthCodeURL()
	if err != nil {
		log.Fatalf("Can not generate auth code url: %v", err)
	}

	return url
}

func (as AuthService) SignUp(input domain.OAuthSignUpInput) (string, error) {
	data, err := as.auth.ExchangeMetaDataByCode(input.Code)
	if err != nil {
		log.Printf("Can not exchange meta data by code: %v", err)
		return "", err
	}

	userProps := domain.UserEntity{
		Name:  strings.Split(data.Email, "@")[0],
		Email: data.Email,
	}

	u, err := as.userRepo.Create(userProps)
	if err != nil {
		log.Printf("Can not create user properly: %v", err)
		return "", err
	}

	userProviderProps := domain.UserProviderEntity{
		Name:   domain.UserProviderNamePrimary,
		UserID: u.ID,
	}

	var _ domain.UserProviderEntity
	_, err = as.userProviderRepo.Create(userProviderProps)
	if err != nil {
		log.Printf("Can not create user provider properly: %w", err)
		return "", err
	}

	return data.AccessToken, nil
}

func (as AuthService) Login() {
}

func NewAuthService(
	userRepo domain.UserRepository,
	userProviderRepo domain.UserProviderRepository,
	auth authenticator.Authenticator,
) AuthService {
	return AuthService{
		userRepo,
		userProviderRepo,
		auth,
	}
}
