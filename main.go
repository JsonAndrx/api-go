package main

import (
	"api-rest/helper"
	"api-rest/config/database"
	"api-rest/users/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Starting server...")
	database.ConectDb()
	routes := gin.Default()
	routes.GET("/test", func(response *gin.Context) {
		users := repository.GetAllUsers()
		response.JSON(http.StatusOK, users)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
