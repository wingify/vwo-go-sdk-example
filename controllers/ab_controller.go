package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wingify/vwo-go-sdk/config"
	"github.com/wingify/vwo-go-sdk/models"
	"github.com/wingify/vwo-go-sdk/util"
	"github.com/gin-gonic/gin"
)

// ABController function uses the configuration values and VWO instance to show usage of Activate API
// along with tracking and displayes the html output
func ABController(c *gin.Context) {
	config := config.GetConfig()
	userID := c.Query("userId")
	if userID == "" {
		userID = util.GetRandomUser()
	}
	campaignKey := c.Query("cKey")
	if campaignKey == "" {
		campaignKey = config.GetString("abCampaignKey")
	}
	abCampaigngoalIdentifier := config.GetString("abCampaignGoalIdentifier")

	instance := models.GetVWOInstance()

	options := make(map[string]interface{})
	options["revenueValue"] = config.GetString("revenueValue")
	options["customVariables"] = make(map[string]interface{})
	options["goalTypeToTrack"] = nil
	options["shouldTrackReturningUser"] = nil

	isPartOfCampaign := false
	variationName := instance.Activate(campaignKey, userID, options)
	if variationName != "" {
		isPartOfCampaign = true
	}

	track := instance.Track(nil, userID, abCampaigngoalIdentifier, options)

	settingsFile, err := json.Marshal(instance.SettingsFile)
	if err != nil {
		fmt.Println(err)
	}

	var class string
	if variationName == "Control" {
		class = "v1"
	} else if variationName == "Variation-1" {
		class = "v2"
	} else {
		class = "v3"
	}

	c.HTML(http.StatusOK, "ab.html", gin.H{
		"campaignType":             "VISUAL_AB",
		"settingsFile":             string(settingsFile),
		"abCampaignKey":            campaignKey,
		"abCampaignGoalIdentifier": abCampaigngoalIdentifier,
		"customVariables":          nil,
		"userID":                   userID,
		"isPartOfCampaign":         isPartOfCampaign,
		"variationName":            variationName,
		"buttonClass":              class,
		"track":                    track,
	})
}
