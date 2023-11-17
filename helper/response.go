package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	meta := Meta{
		Message: message,
		Status:  "success",
		Code:    http.StatusOK,
	}
	jsonRes := Response{
		Meta: meta,
		Data: data,
	}
	c.JSON(http.StatusOK, jsonRes)
}

func BadRequestResponse(c *gin.Context, message string, data interface{}) {
	meta := Meta{
		Message: message,
		Status:  "fail",
		Code:    http.StatusBadRequest,
	}
	jsonRes := Response{
		Meta: meta,
		Data: data,
	}
	c.JSON(http.StatusBadRequest, jsonRes)
}
