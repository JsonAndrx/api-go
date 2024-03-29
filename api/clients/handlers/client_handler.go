package handlers

import (
	"github.com/gin-gonic/gin"
	"api-rest/api/clients/services"
)

func RoutesClient(router *gin.RouterGroup) {
	routeClient := router.Group("/clients")
	{
		routeClient.POST("/uploadclients/", UploadClientHandler)
	}
}

func UploadClientHandler(c *gin.Context) {
	services.UploadClientSerive(c)
}