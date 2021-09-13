package app

import (

	// gin import

	//import local package
	"go-rest-boilerplate/src/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
    app := gin.New()

	routes.RoutesSetup(app)

    app.Run(":5050")
}