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
// @contact.name Cozy
// @contact.url https://github.com/MasDeny
// @license.name MIT
// @license.url https://github.com/MasDeny/go-rest-boilerplate
// @BasePath /api/v1
// @host api.example.a2os.club
func Run() {
    app := gin.New()

	routes.RoutesSetup(app)

    app.Run(":5050")
}