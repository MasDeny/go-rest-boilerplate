package controllers

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
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "Undermaintenance !"
// @Router /ping [get]
func ServerHealth(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
			"status": "OK",
	})
}

func WelcomeGreeting(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome To Go Rest API",
		"date": time.Now().UTC().Format(time.RFC1123Z),
	})
}