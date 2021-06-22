package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wingify.com/vwo-go-sdk-example/config"
	"wingify.com/vwo-go-sdk-example/models"
	"wingify.com/vwo-go-sdk-example/util"
	"github.com/gin-gonic/gin"
)

// PushController function gets the configuration values and VWO instance to
// check the Push API and displayes the html output
func PushController(c *gin.Context) {
	config := config.GetConfig()
	userID := c.Query("userId")
	tagKey := c.Query("tagKey")
	tagValue := c.Query("tagValue")
	if userID == "" {
		userID = util.GetRandomUser()
	}

	if tagKey == "" {
		tagKey = config.GetString("tagKey")
	}
	if tagValue == "" {
		tagValue = config.GetString("tagValue")
	}

	instance := models.GetVWOInstance()

	result := instance.Push(tagKey, tagValue, userID)

	settingsFile, err := json.Marshal(instance.SettingsFile)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "push.html", gin.H{
		"settingsFile": string(settingsFile),
		"tagKey":       tagKey,
		"tagValue":     tagValue,
		"userID":       userID,
		"result":       result,
	})
}
