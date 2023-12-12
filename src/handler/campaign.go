package handler

import (
	"crowdfunding-api/src/domain/campaign"
	"crowdfunding-api/src/domain/constant"
	"crowdfunding-api/src/domain/user"
	"crowdfunding-api/src/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler() *campaignHandler {
	service := campaign.NewService()
	return &campaignHandler{service: service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaignData, err := h.service.GetCampaigns(userID)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	helper.SuccessResponse(c, "GET campaigns", campaign.FormatGetListCampaignResponse(campaignData))
}

func (h *campaignHandler) GetCampaignDetail(c *gin.Context) {
	var input campaign.GetByIdInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	campaignData, err := h.service.GetCampaignDetail(input.ID)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	helper.SuccessResponse(c, "GET detail campaign", campaign.FormatCampaignDetailResponse(campaignData))
	// return
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	currentUser := c.MustGet("currentUser")

	input.User = currentUser.(user.User)

	err = h.service.CreateCampaign(input)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	helper.SuccessResponse(c, "CREATE campaign", nil)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var input campaign.UpdateCampaignInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	err = c.ShouldBindJSON(&input.Data)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	err = h.service.UpdateCampaign(input)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	helper.SuccessResponse(c, "UPDATE campaign", nil)
}
