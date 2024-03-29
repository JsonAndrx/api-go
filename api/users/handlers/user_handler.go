package handlers

import (
	"api-rest/api/users/services"

	"github.com/gin-gonic/gin"
)

func RoutesUser(router *gin.RouterGroup) {
	routeUser := router.Group("/users")
	{
		routeUser.GET("/", services.GetUsersService)
		routeUser.POST("/create/", CreateUserHandler)
	}
}


func CreateUserHandler(c *gin.Context) {
	services.CreateUserService(c)
}
