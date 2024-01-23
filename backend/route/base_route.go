package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type BaseRoute struct {
	router *gin.Engine
}

func (br BaseRoute) Setup() {
	br.router.GET("/health", func(c *gin.Context) {
		fmt.Println("in health")
		c.JSON(200, gin.H{
			"message": "Alive...",
		})
	})
}

func NewBaseRoute(
	router *gin.Engine,
) BaseRoute {
	return BaseRoute{
		router: router,
	}
}
