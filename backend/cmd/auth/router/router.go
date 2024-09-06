package router

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/cmd/auth/docs"
	"github.com/Dominic0512/serverless-auth-boilerplate/route"
	"github.com/google/wire"
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRouter(
	baseRoute route.BaseRoute,
	authRoute route.AuthRoute,
) Routes {
	return Routes{
		baseRoute,
		authRoute,
	}
}

func (r Routes) Setup() {
	docs.SwaggerInfo.BasePath = "/api/v1"
	for _, route := range r {
		route.Setup()
	}
}

var ProviderSet = wire.NewSet(NewRouter)
