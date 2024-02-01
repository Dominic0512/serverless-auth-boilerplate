package runner

import (
	"github.com/gin-gonic/gin"
)

type GinRunner struct {
	engine *gin.Engine
}

func NewGinRunner(g *gin.Engine) *GinRunner {
	return &GinRunner{
		engine: g,
	}
}

func (gr GinRunner) Run() {
	gr.engine.Run()
}
