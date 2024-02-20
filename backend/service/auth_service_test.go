package service

import (
	"testing"

	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/authenticator"
	"github.com/Dominic0512/serverless-auth-boilerplate/mocks"
	"github.com/Dominic0512/serverless-auth-boilerplate/pkg/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_GenerateAuthURL(t *testing.T) {
	repo := new(mocks.UserRepository)
	auth := new(mocks.Authenticator)

	auth.On("GenerateAuthCodeURL").Return(
		func() string {
			return ""
		},
		func() error {
			return nil
		},
	)

	as := NewAuthService(repo, auth)

	t.Run("it should return oauth url without err", func(t *testing.T) {
		_, err := as.auth.GenerateAuthCodeURL()

		a := assert.New(t)
		a.Nil(err)
	})
}

func TestAuthService_SignUp(t *testing.T) {
	dummyJWTToken := "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwiaXNzIjoiaHR0cHM6Ly9kZXYtd3A4b2h0dXFobXJvNW13YS51cy5hdXRoMC5jb20vIn0..fFxOTGtxT9EcM621.RuzzHt-ToBnPbD8mjnSqg6F9HglV2oeaQPToQPDcTGwUa_3DyPrM_2mCZ1sZxxndnqrsQG7RW_Nv1gBHuEYm0i2V8lTx2YS0GbI5YeYe_KS5STeHRqgQpPIGiIj3TgKLSIzE4UznBy_3vOIkFbnsj6bbJRVu9YV9OIiQUkxRlx38isZotjHC91m-XJj02-0ZOcgnBMLQQSLix0Ti-w332J15ViD85Ps3E6d26RqIiyaSzd2PR1kxb5ejmw2WC2VXgZBNBTmu3ZytiLunDYydK-HqxPrujLptU6-utqcCpo_UaPVNH2ahoSCDZiOyMdWnBtbFiCDEFRNKsmdBU6YU83tX.LLQgf8yZ6NiMHasatqanHw"
	repo := new(mocks.UserRepository)
	auth := new(mocks.Authenticator)

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

	as := NewAuthService(repo, auth)

	t.Run("it will return access token", func(t *testing.T) {
		input := domain.OAuthSignUpInput{
			Code: "",
		}
		token, err := as.SignUp(input)

		validator := validate.NewValidator()
		vErr := validator.Validate.Var(token, "required,jwt")

		a := assert.New(t)
		a.Nil(vErr)
		a.Nil(err)
	})
}
