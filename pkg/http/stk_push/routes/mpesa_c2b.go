package routes

import (
	"github.com/kisese/golang_mpesa/pkg/http/stk_push/controllers"
	"github.com/kisese/golang_mpesa/pkg/infrastructure"
)

type STKRoute struct {
	controller controllers.STKController
	Handler    infrastructure.GinRouter
}

func NewSTKRoute(
	controller controllers.STKController,
	handler infrastructure.GinRouter,

) STKRoute {
	return STKRoute{
		controller: controller,
		Handler:    handler,
	}
}

func (location STKRoute) Setup() {
	router := location.Handler.Gin.Group("/stk") //Router group
	{
		router.POST("/request", location.controller.ProcessSTKPushRequest)
		router.POST("/callback", location.controller.ProcessSTKCallback)
	}
}
