package handler

import (
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
		helper.BadRequestResponse(c, "Create user gagal", err.Error())
		return
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		helper.BadRequestResponse(c, "Create user gagal", err.Error())
		return
	}

	helper.SuccessResponse(c, "Create user success", user)
}
