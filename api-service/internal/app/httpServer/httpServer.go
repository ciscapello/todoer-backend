package httpServer

import (

	// "api-service/internal/app/http/handlers/auth"

	"github.com/ciscapello/api-service/internal/app/httpServer/handlers"
	"github.com/ciscapello/api-service/internal/services"
	"github.com/ciscapello/api-service/pkg/tokenManager"

	"github.com/gin-gonic/gin"
)

func RunServer(port string, services *services.Services, tokenManager *tokenManager.TokenManager) {
	r := gin.Default()

	handler := handlers.NewHandler(services, tokenManager, r)

	handler.RunHandler()

	r.Run()
	// r.Use(middleware.Middleware)

	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
