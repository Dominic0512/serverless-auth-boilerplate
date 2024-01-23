package route

import (
	"github.com/google/wire"
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoute(
	baseRoute BaseRoute,
	authRoute AuthRoute,
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

var ProviderSet = wire.NewSet(NewBaseRoute, NewAuthRoute, NewRoute)
