package framework

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewGinFramework)
