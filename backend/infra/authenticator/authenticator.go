package authenticator

import (
	"github.com/google/wire"
)

type AuthMetaData struct {
	AccessToken   string
	Email         string
	EmailVerified bool
}

type Authenticator interface {
	GenerateAuthCodeURL() (string, error)
	ExchangeMetaDataByCode(code string) (*AuthMetaData, error)
}

var ProviderSet = wire.NewSet(
	NewAuth0Authenticator,
	wire.Bind(new(Authenticator), new(*Auth0Authenticator)),
)
