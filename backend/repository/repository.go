package repository

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserRepository,
	wire.Bind(new(domain.UserRepository), new(*UserRepository)))
