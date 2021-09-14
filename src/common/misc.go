package common

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//
// @Summary Test Api Server
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Success 200 {object} Response
// @Router /ping [get]
func ServerHealth(c *gin.Context){
	Error400(c, "Server Maintenance", 5003)
	// c.JSON(http.StatusOK, BaseResponse("Status Server OK"))
}

func WelcomeGreeting(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome To Go Rest API",
		"date": time.Now().UTC().Format(time.RFC1123Z),
	})
}