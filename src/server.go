package app

import (

	// gin import

	//import local package
	"go-rest-boilerplate/src/routes"

	"github.com/gin-gonic/gin"
)

// Run starts the app
// @title GoGin Boilerplate API
// @version 1.0
// @description This is my go gin boilerplate api server.
// @termsOfService http://swagger.io/terms/
// @contact.name MasDeny
// @contact.url https://github.com/MasDeny
// @license.name MIT
// @license.url https://github.com/MasDeny/go-rest-boilerplate
// @host localhost:5050
// @BasePath /api/v1
func Run() {
    app := gin.New()

	routes.RoutesSetup(app)

    app.Run(":5050")
}