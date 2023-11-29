package helper

import (
	"crowdfunding-api/constant"
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

func ErrorResponse(c *gin.Context, exception constant.Exception, data interface{}) {
	meta := Meta{
		Message: exception.Message,
		Status:  "fail",
		Code:    exception.StatusCode,
	}
	jsonRes := Response{
		Meta: meta,
		Data: data,
	}
	c.JSON(http.StatusBadRequest, jsonRes)
}

func AbortResponse(c *gin.Context, exception constant.Exception, data interface{}) {
	meta := Meta{
		Message: exception.Message,
		Status:  "fail",
		Code:    exception.StatusCode,
	}
	jsonRes := Response{
		Meta: meta,
		Data: data,
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, jsonRes)
}
