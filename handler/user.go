package handler

import (
	"crowdfunding-api/constant"
	"crowdfunding-api/helper"
	"crowdfunding-api/user"
	"fmt"

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

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	resp, err := h.userService.CheckEmail(input)
	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	var message string
	if resp == false {
		message = "Email not available"
	} else {
		message = "Email available"
	}

	helper.SuccessResponse(c, message, resp)
}

func (h *userHandler) UpdateAvatar(c *gin.Context) {

	file, err := c.FormFile("avatar")

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	_, err = h.userService.UpdateAvatar(userID, path)

	if err != nil {
		helper.ErrorResponse(c, constant.BadRequest, err.Error())
		return
	}

	data := gin.H{"is_uploaded": true}

	helper.SuccessResponse(c, "Success Upload Avatar", data)
}
