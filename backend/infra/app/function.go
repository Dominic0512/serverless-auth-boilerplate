package app

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/runner"
	"github.com/Dominic0512/serverless-auth-boilerplate/route"
)

type FunctionApp struct {
	runner *runner.LambdaRunner
}

func NewFunctionApp(
	routes route.Routes,
	runner *runner.LambdaRunner,
) (*FunctionApp, error) {
	routes.Setup()
	return &FunctionApp{
		runner: runner,
	}, nil
}

func (fa *FunctionApp) Start() {
	fa.runner.Run()
}
