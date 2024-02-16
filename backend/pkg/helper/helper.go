package helper

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewBcryptPasswordHelper,
	wire.Bind(new(PasswordHelper), new(*BcryptPasswordHelper)),
)
