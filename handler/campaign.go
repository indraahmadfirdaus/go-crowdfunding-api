package handler

import (
	"crowdfunding-api/campaign"
	"crowdfunding-api/constant"
	"crowdfunding-api/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaignData, err := h.service.GetCampaigns(userID)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, nil)
		return
	}

	helper.SuccessResponse(c, "GET campaigns", campaign.FormatGetListCampaignResponse(campaignData))
}

func (h *campaignHandler) GetCampaignDetail(c *gin.Context) {
	var input campaign.GetByIdInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, nil)
		return
	}

	campaignData, err := h.service.GetCampaignDetail(input.ID)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, nil)
		return
	}

	helper.SuccessResponse(c, "GET detail campaign", campaign.FormatCampaignDetailResponse(campaignData))
	// return
}
