package main

import (
	"github.com/kisese/golang_mpesa/pkg/http/stk/controllers"
	"github.com/kisese/golang_mpesa/pkg/http/stk/routes"
	"github.com/kisese/golang_mpesa/pkg/http/stk/service"
	"github.com/kisese/golang_mpesa/pkg/infrastructure"
)

func init() {
	infrastructure.LoadEnv()
	infrastructure.InitLogger()
}

func main() {
	router := infrastructure.NewGinRouter()
	stkService := service.NewStkRequestService()
	stkController := controllers.NewSTRRequestController(stkService)
	stkRoute := routes.NewSTKRoute(stkController, router)
	stkRoute.Setup()

	router.Gin.Run(":8080") //server started on 8000 port
}
