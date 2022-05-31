package infrastructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinRouter struct {
	Gin *gin.Engine
}

func NewGinRouter() GinRouter {
	httpRouter := gin.Default()

	httpRouter.GET("/", func(c *gin.Context) {

		Log.Infow("GET Index",
			"url", "/",
			"result_code", http.StatusOK,
			"data", "Up and Running...",
		)

		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	return GinRouter{
		Gin: httpRouter,
	}
}
