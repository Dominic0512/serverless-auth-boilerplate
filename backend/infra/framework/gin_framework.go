package framework

import (
	"github.com/gin-gonic/gin"
)

func NewGinFramework() *gin.Engine {
	g := gin.New()
	return g
}
