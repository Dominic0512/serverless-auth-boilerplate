//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/controller"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/app"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/config"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/router"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/runner"
	"github.com/Dominic0512/serverless-auth-boilerplate/pkg/validate"
	"github.com/Dominic0512/serverless-auth-boilerplate/repository"
	"github.com/Dominic0512/serverless-auth-boilerplate/route"
	"github.com/Dominic0512/serverless-auth-boilerplate/service"
	"github.com/google/wire"
)

func InitializeApp() (*app.FunctionApp, error) {
	wire.Build(
		config.ProviderSet,
		database.ProviderSet,
		router.ProviderSet,
		repository.ProviderSet,
		service.ProviderSet,
		controller.ProviderSet,
		route.ProviderSet,
		validate.ProviderSet,
		runner.ProviderSet,
		app.ProviderSet,
	)
	return &app.FunctionApp{}, nil
}
