package repository

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserRepository,
	NewUserProviderRepository,
	wire.Bind(new(domain.UserRepository), new(*UserRepository)),
	wire.Bind(new(domain.UserProviderRepository), new(*UserProviderRepository)),
)
