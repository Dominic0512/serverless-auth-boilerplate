package runner

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewLambdaRunner, NewGin)