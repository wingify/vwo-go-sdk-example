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

// FeatureRolloutController function uses the configuration values and VWO instance to check
// whether a feature is enabled or not for the given user and displayes the html output
func FeatureRolloutController(c *gin.Context) {
	config := config.GetConfig()
	userID := c.Query("userId")
	if userID == "" {
		userID = util.GetRandomUser()
	}
	campaignKey := c.Query("cKey")
	if campaignKey == "" {
		campaignKey = config.GetString("featureRolloutCampaignKey")
	}

	instance := models.GetVWOInstance()

	isEnabled := instance.IsFeatureEnabled(campaignKey, userID, nil)

	settingsFile, err := json.Marshal(instance.SettingsFile)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "feature_rollout.html", gin.H{
		"campaignType":                        "FEATURE_ROLLOUT",
		"settingsFile":                        string(settingsFile),
		"campaifeatureRolloutCampaignKeyType": campaignKey,
		"customVariables":                     nil,
		"userID":                              userID,
		"isUserPartOfFeatureRolloutCampaign":  isEnabled,
	})

}
