package controllers

import (
	"github.com/gin-gonic/gin"
	"wingify.com/vwo-go-sdk-example/models"
)

// Manual Flush for event batching
func BatchController(c *gin.Context) {

	instance := models.GetVWOInstance()
	instance.FlushEvents()
}
