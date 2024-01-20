package app

import (
	"database/sql"
	"net/http"

	"github.com/ciscapello/api-service/internal/app/database"
	"github.com/ciscapello/api-service/internal/app/httpServer"
	"github.com/ciscapello/api-service/internal/config"
	"github.com/ciscapello/api-service/internal/services"
	"github.com/ciscapello/api-service/pkg/tokenManager"
)

type App struct {
	httpServer http.Server
	database   sql.DB
}

func Init() *App {
	return &App{}
}

func (app *App) Run() {
	conf := config.New()

	// TODO: Run Database
	pool := database.Connect()

	services := services.Init(pool)

	tokenManager := tokenManager.NewTokenManager()

	// userService := services.NewUserService(pool)
	// authService := services.NewAuthService()

	// // TODO: Run HTTP server

	httpServer.RunServer(conf.Port, services, tokenManager)

	// httpServer.RunServer(conf.Port, services, tokenManager)

}
