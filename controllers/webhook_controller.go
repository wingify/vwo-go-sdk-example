package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wingify.com/vwo-go-sdk-example/config"
	"wingify.com/vwo-go-sdk-example/models"
)

// PushController function gets the configuration values and VWO instance to
// check the Push API and displayes the html output
func WebhookController(c *gin.Context) {
	config := config.GetConfig()
	instance := models.GetVWOInstance()
	signature, exists := c.Request.Header["X-Vwo-Auth"]

	if config.IsSet("webhookSecret") && exists {
		webhookAuthKey := config.GetString("webhookSecret")
		if signature[0] != webhookAuthKey {
			fmt.Println("VWO webhook authentication failed. Please check.")
			return
		} else {
			fmt.Println("VWO webhook authenticated successfully.")
		}
	} else {
		fmt.Println("Skipping Webhook Authentication as webhookAuthKey is not provided in config")
	}
	instance.GetAndUpdateSettingsFile()
}
