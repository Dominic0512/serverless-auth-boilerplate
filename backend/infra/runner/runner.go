package runner

import "github.com/google/wire"

type Runner interface {
	Run()
}

var ProviderSet = wire.NewSet(
	NewLambdaRunner,
	NewGinRunner,
)
