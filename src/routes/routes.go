package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
func RoutesSetup(app *gin.Engine){
	api := app.Group("/api")
	v1 := api.Group("/v1")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
    })

	v1.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome To Go Rest API",
			"date": time.Now().Day(),
		})
	})
}