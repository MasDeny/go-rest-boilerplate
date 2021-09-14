package routes

import (
	"go-rest-boilerplate/src/common"
	_ "go-rest-boilerplate/src/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)
func RoutesSetup(app *gin.Engine){

	// Error Handle on Not Found
	app.NoRoute(func(c *gin.Context) {
		common.Error404(c)
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1.GET("/ping", common.ServerHealth)

	v1.GET("/", common.WelcomeGreeting)
}