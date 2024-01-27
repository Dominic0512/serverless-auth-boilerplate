package runner

import (
	"github.com/gin-gonic/gin"
)

func NewGin() *gin.Engine {
	g := gin.New()
	return g
}
