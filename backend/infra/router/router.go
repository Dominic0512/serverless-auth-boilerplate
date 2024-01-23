package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func NewRouter() *gin.Engine {
	g := gin.New()
	return g
}

var ProviderSet = wire.NewSet(NewRouter)
