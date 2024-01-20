package main

import (
	"fmt"

	_ "github.com/ciscapello/api-service/internal/app/database/migrations"

	"github.com/ciscapello/api-service/internal/app"
)

func main() {
	fmt.Println("Application is runnin")
	// TODO: Get config
	app := app.Init()
	app.Run()

	// conf := config.New()

	// // TODO: Run Database
	// pool := database.Connect()

	// userService := services.NewUserService(pool)
	// authService := services.NewAuthService()

	// // // TODO: Run HTTP server
	// http.RunServer(conf.Port, userService, authService)

	// TODO: Run gRPC server

	// TODO: Handle graceful shutdown
}
