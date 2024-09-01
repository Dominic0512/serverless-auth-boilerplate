package router

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/cmd/user/docs"
	"github.com/Dominic0512/serverless-auth-boilerplate/route"
	"github.com/google/wire"
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRouter(
	baseRoute route.BaseRoute,
	userRoute route.UserRoute,
) Routes {
	return Routes{
		baseRoute,
		userRoute,
	}
}

func (r Routes) Setup() {
	docs.SwaggerInfo.BasePath = "/api/users"
	for _, route := range r {
		route.Setup()
	}
}

var ProviderSet = wire.NewSet(NewRouter)
