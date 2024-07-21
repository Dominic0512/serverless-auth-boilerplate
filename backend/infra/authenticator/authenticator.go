package authenticator

import (
	"github.com/google/wire"
)

type AuthMetaData struct {
	AccessToken   string
	Email         string
	EmailVerified bool
	Picture       string
	Sub           string
	Provider      string
}

type Authenticator interface {
	GenerateAuthCodeURL() (string, error)
	ExchangeMetaDataByCode(code string) (*AuthMetaData, error)
	TransformProviderName(name string) (*string, error)
}

var ProviderSet = wire.NewSet(
	NewAuth0Authenticator,
	wire.Bind(new(Authenticator), new(*Auth0Authenticator)),
)
