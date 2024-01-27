package router

import (
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
	for _, route := range r {
		route.Setup()
	}
}

var ProviderSet = wire.NewSet(NewRouter)
