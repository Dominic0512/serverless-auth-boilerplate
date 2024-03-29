// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/cmd/user/gin/app"
	"github.com/Dominic0512/serverless-auth-boilerplate/cmd/user/router"
	"github.com/Dominic0512/serverless-auth-boilerplate/controller"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/config"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/framework"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/runner"
	"github.com/Dominic0512/serverless-auth-boilerplate/pkg/helper"
	"github.com/Dominic0512/serverless-auth-boilerplate/pkg/validate"
	"github.com/Dominic0512/serverless-auth-boilerplate/repository"
	"github.com/Dominic0512/serverless-auth-boilerplate/route"
	"github.com/Dominic0512/serverless-auth-boilerplate/service"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func InitializeApp() (*app.App, error) {
	engine := framework.NewGinFramework()
	baseRoute := route.NewBaseRoute(engine)
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	databaseDatabase, err := database.NewDatabase(configConfig)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(databaseDatabase)
	bcryptPasswordHelper := helper.NewBcryptPasswordHelper()
	userService := service.NewUserService(userRepository, bcryptPasswordHelper)
	validator := validate.NewValidator()
	userController := controller.NewUserController(userService, validator)
	userRoute := route.NewUserRoute(engine, userController)
	routes := router.NewRouter(baseRoute, userRoute)
	ginRunner := runner.NewGinRunner(engine)
	appApp, err := app.NewApp(routes, ginRunner)
	if err != nil {
		return nil, err
	}
	return appApp, nil
}
