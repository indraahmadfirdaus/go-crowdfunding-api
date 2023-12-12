package routes

import (
	"crowdfunding-api/src/handler"
	"crowdfunding-api/src/middleware"

	"github.com/gin-gonic/gin"
)

var userHandler = handler.NewUserHandler()

func (r routes) InitUserRoute(rg *gin.RouterGroup) {
	api := rg.Group("/users")
	m := middleware.NewMiddleware()

	api.POST("/", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/check-email", userHandler.CheckEmailAvailability)
	api.POST("/upload-avatar", m.AuthMiddleware(), userHandler.UpdateAvatar)
}
