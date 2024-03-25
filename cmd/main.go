package main

import (
	// "fmt"

	"api-rest/api/users/handlers"
	"api-rest/api/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"net/http"
)

func main() {
	log.Info().Msg("Starting server...")
	routes := gin.Default()

	routes_api := routes.Group("/api/v1")
	handlers.RoutesUser(routes_api)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	response.ErrorResponse("", err.Error())
}
