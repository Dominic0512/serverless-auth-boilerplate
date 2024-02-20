package service

import (
	"log"
	"strings"

	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/authenticator"
)

type AuthService struct {
	repo domain.UserRepository
	auth authenticator.Authenticator
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

	var _ string
	_, err = as.repo.Create(userProps)
	if err != nil {
		log.Printf("Can not create user properly: %v", err)
		return "", err
	}

	return data.AccessToken, nil
}

func (as AuthService) Login() {
}

func NewAuthService(repo domain.UserRepository, auth authenticator.Authenticator) AuthService {
	return AuthService{
		repo,
		auth,
	}
}
