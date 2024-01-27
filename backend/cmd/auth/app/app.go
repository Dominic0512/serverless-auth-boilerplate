package app

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/cmd/auth/router"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/runner"
	"github.com/google/wire"
)

type App struct {
	runner *runner.LambdaRunner
}

func NewApp(
	routes router.Routes,
	runner *runner.LambdaRunner,
) (*App, error) {
	routes.Setup()
	return &App{
		runner: runner,
	}, nil
}

func (fa *App) Start() {
	fa.runner.Run()
}

var ProviderSet = wire.NewSet(NewApp)
