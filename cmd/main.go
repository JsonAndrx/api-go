package main

import (
	// "fmt"

	"api-rest/api/users/handlers"
	"api-rest/api/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"api-rest/config/database"

	"net/http"
)

func main() {
	log.Info().Msg("Starting server...")
	database.ConectDb()

	routes := gin.Default()

	routesApi := routes.Group("/api/v1")
	handlers.RoutesUser(routesApi)

	server := &http.Server{
		Addr:    ":8081",
		Handler: routes,
	}

	err := server.ListenAndServe()
	response.ErrorResponse("", err.Error())
}
