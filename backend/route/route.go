package route

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewBaseRoute,
	NewAuthRoute,
	NewUserRoute,
)
