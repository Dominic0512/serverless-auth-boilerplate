package helper

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewBcryptPasswordHelper,
	NewEntTxHelper,
	wire.Bind(new(PasswordHelper), new(*BcryptPasswordHelper)),
	wire.Bind(new(TxHelper), new(*EntTxHelper)),
)
