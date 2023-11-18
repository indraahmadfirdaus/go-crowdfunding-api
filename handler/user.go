package handler

import (
	"crowdfunding-api/constant"
	"crowdfunding-api/helper"
	"crowdfunding-api/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	helper.SuccessResponse(c, "Create user success", user)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorResponse(c, constant.LoginFailed, err.Error())
		return
	}

	user, err := h.userService.Login(input)
	if err != nil {
		helper.ErrorResponse(c, constant.LoginFailed, err.Error())
		return
	}

	helper.SuccessResponse(c, "Successfully Login", user)
}
