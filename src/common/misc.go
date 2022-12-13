package common

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// @Tags Initialize
// @Summary Test Api Server
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Success 200 {object} Response
// @Router /ping [get]
func ServerHealth(c *gin.Context) {
	// Error400(c, "Server Maintenance", 5003)
	log.Println(`don't scare !!`)
	c.JSON(http.StatusOK, BaseResponse("Status Server OK"))
}

// @Tags Initialize
// @Summary Server Greeting
// @Produce json
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router / [get]
func WelcomeGreeting(c *gin.Context) {
	res := gin.H{
		"message": "Welcome To Go Rest API",
		"date":    time.Now().UTC().Format(time.RFC1123Z),
	}
	c.JSON(http.StatusOK, BaseResponse(res))
}
