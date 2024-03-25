package handlers

import (
	"api-rest/api/users/services"

	"github.com/gin-gonic/gin"
)

func RoutesUser(router *gin.RouterGroup) {
	routes_user := router.Group("/users")
	{
		routes_user.GET("/", services.GetUsersService)
		routes_user.POST("/create/", services.CreateUserService)
	}
}
