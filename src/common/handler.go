package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status bool `json:"status"`
	Result interface{} `json:"result"`
	Status_code int `json:"status_code"`
}

type ErrorResponse struct {
	Status bool `json:"status"`
	Result string `json:"result"`
	Status_code int `json:"status_code"`
}

func BaseResponse(result interface{}) *Response{
	return &Response{
		Status: true,
		Result: result,
		Status_code: 0,
	}
}

func ErrorHandlerResponse(message string, code int)*ErrorResponse{
	return &ErrorResponse{
		Status: false,
		Result: message,
		Status_code: code,
	}
}

func Error400(c *gin.Context, message string, code int) {
	c.JSON(http.StatusBadRequest, ErrorHandlerResponse(message, code))
}

func Error500(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, ErrorHandlerResponse(message, 5000))
}

func Error401(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, ErrorHandlerResponse(message, 4001))
}

func Error404(c *gin.Context) {
	c.JSON(http.StatusNotFound, ErrorHandlerResponse("Not Found", 4004))
}

