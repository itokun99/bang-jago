package router

import (
	"github.com/itokun99/bang-jago/app/config"
	"github.com/itokun99/bang-jago/app/service"

	"github.com/gin-gonic/gin"
)

var appConfig = config.GetConfig()

// function for get router engine instance
func GetRouter() *gin.Engine {

	// setup environtment to production
	if appConfig.Server.ServerEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	return router
}

// function for init router in app
func InitRouter() {
	app := GetRouter()

	// middleware
	app.Use(service.CORS())

	// endpoint list

	// Service Get Base Url
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"code":    200,
			"message": "Bang Jago Boilerplate for REST API",
		})
	})

	app.Run(":" + appConfig.Server.ServerPort)
}
