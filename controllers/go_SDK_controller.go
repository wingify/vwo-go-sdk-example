package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wingify/vwo-go-sdk/models"
	"github.com/gin-gonic/gin"
)

// GoSDKController function displayes the settings file
func GoSDKController(c *gin.Context) {
	instance := models.GetVWOInstance()

	settingsFile, err := json.Marshal(instance.SettingsFile)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "go_SDK.html", gin.H{
		"settingsFile": string(settingsFile),
	})
}
