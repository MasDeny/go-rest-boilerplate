package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status      bool        `json:"status"`
	Result      interface{} `json:"result"`
	Status_code int         `json:"status_code"`
}

type ErrorResponse struct {
	Status      bool   `json:"status"`
	Result      string `json:"result"`
	Status_code int    `json:"status_code"`
}

func BaseResponse(result interface{}) *Response {
	return &Response{
		Status:      true,
		Result:      result,
		Status_code: 0,
	}
}

func ErrorHandlerResponse(message string, code int) *ErrorResponse {
	return &ErrorResponse{
		Status:      false,
		Result:      message,
		Status_code: code,
	}
}

func BadRequestException(c *gin.Context, message string, code int) {
	c.JSON(http.StatusBadRequest, ErrorHandlerResponse(message, code))
}

func InternalServerErrorException(c *gin.Context, message string) {
	c.JSON(
		http.StatusInternalServerError,
		ErrorHandlerResponse(`Internal Server Error`, 5000),
	)
}

func UnauthorizedException(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, ErrorHandlerResponse(message, 4001))
}

func NotFoundException(c *gin.Context) {
	c.JSON(http.StatusNotFound, ErrorHandlerResponse("Not Found", 4004))
}

func ForbiddenException(c *gin.Context) {
	c.JSON(http.StatusForbidden, ErrorHandlerResponse("", 4003))
}

func UnsupportedException(c *gin.Context) {
	c.JSON(
		http.StatusUnsupportedMediaType,
		ErrorHandlerResponse("Unsupported Media Type", 4015),
	)
}

func BadGatewayException(c *gin.Context) {
	c.JSON(
		http.StatusBadGateway,
		ErrorHandlerResponse("Bad Gateway", 5002),
	)
}

func ServiceUnavailableException(c *gin.Context) {
	c.JSON(
		http.StatusServiceUnavailable,
		ErrorHandlerResponse("Service Unavailable", 5003),
	)
}
