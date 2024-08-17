package authenticator

import (
	"github.com/google/wire"
	"golang.org/x/oauth2"
)

type AuthMetaData struct {
	Token         *oauth2.Token
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
