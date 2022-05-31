package main

import (
	"github.com/kisese/golang_mpesa/pkg/http/stk_push/controllers"
	"github.com/kisese/golang_mpesa/pkg/http/stk_push/routes"
	"github.com/kisese/golang_mpesa/pkg/infrastructure"
	"github.com/kisese/golang_mpesa/pkg/queue"
	"github.com/urfave/cli"
	"os"
)

var (
	app *cli.App
)

func init() {
	infrastructure.LoadEnv()
	infrastructure.InitLogger()

	app = cli.NewApp()

	app.Name = "M-Pesa C2B App (STK Push, USSD Push)"
	app.Usage = "An app that shows how to consume the MPESA C2B APIs for STK Push/USSD Push powered by golang, rabbitmq"
	app.Author = "Brian Kisese"
	app.Email = "brayokisese@gmail.com"
}

func main() {

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run the main app that takes has the main API",
			Action: func(c *cli.Context) {
				infrastructure.Log.Infow("main app")
				startMainApp()
			},
		},
		{
			Name:  "worker",
			Usage: "Run the worker that will consume tasks from the queue",
			Action: func(c *cli.Context) {
				infrastructure.Log.Infow("worker")

				queue.StartQueueConsumer()
			},
		},
	}

	app.Run(os.Args)
}

func startMainApp() {
	router := infrastructure.NewGinRouter()
	stkController := controllers.NewSTRRequestController()
	stkRoute := routes.NewSTKRoute(stkController, router)
	stkRoute.Setup()

	router.Gin.Run(":8080")
}
