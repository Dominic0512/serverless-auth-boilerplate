package controller

import (
	"github.com/google/wire"

	"github.com/Dominic0512/serverless-auth-boilerplate/controller/auth"
	"github.com/Dominic0512/serverless-auth-boilerplate/controller/user"
)

var ProviderSet = wire.NewSet(auth.NewAuthController, user.NewUserController)
