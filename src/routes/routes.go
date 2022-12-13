package routes

import (
	"go-rest-boilerplate/src/common"
	docs "go-rest-boilerplate/src/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title           Swagger Example API Test
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support Test
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func RoutesSetup(app *gin.Engine) {

	// Error Handle on Not Found
	app.NoRoute(func(c *gin.Context) {
		common.NotFoundException(c)
	})

	api := app.Group("/api")

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := api.Group("/v1")
	{
		v1.GET("/", common.WelcomeGreeting)
		v1.GET("/ping", common.ServerHealth)
		v1.GET("/test", common.Test)
		// initialize := v1.Group("/initialize")
		// {
		// }
	}
}
