package routes

import (
	"crowdfunding-api/src/handler"
	"crowdfunding-api/src/middleware"

	"github.com/gin-gonic/gin"
)

var campaignHandler = handler.NewCampaignHandler()

func (r routes) InitCampaignRoute(rg *gin.RouterGroup) {
	api := rg.Group("/campaigns")
	m := middleware.NewMiddleware()

	api.GET("/", campaignHandler.GetCampaigns)
	api.GET("/:id", campaignHandler.GetCampaignDetail)
	api.POST("/", m.AuthMiddleware(), campaignHandler.CreateCampaign)
	api.PUT("/:id", m.AuthMiddleware(), campaignHandler.UpdateCampaign)

}
