package server

import (
	"github.com/wingify/vwo-go-sdk/controllers"
	"github.com/gin-gonic/gin"
)

//NewRouter function controls the routes
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/*")

	ping := new(controllers.StatusController)
	router.GET("/ping", ping.Status)
	router.GET("/ab", controllers.ABController)
	router.GET("/feature-rollout", controllers.FeatureRolloutController)
	router.GET("/feature-test", controllers.FeatureTestController)
	router.GET("/push", controllers.PushController)
	router.GET("/go-sdk", controllers.GoSDKController)
	router.GET("/", controllers.HomePage)

	return router
}
