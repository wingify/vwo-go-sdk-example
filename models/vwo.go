package models

import (
	"fmt"

	"github.com/wingify/vwo-go-sdk/config"
	vwo "github.com/wingify/vwo-go-sdk"
	"github.com/wingify/vwo-go-sdk/pkg/api"
)

// GetVWOInstance function initializes the settings file and launches the
// VWO instance using configuration values and returns the insttance
func GetVWOInstance() *api.VWOInstance {
	config := config.GetConfig()

	// User Storage
	storage := &UserStorageData{}

	// Custom Logger (Google Logger)
	// logs := &LogS{}

	// Instance with userStorage and customLogger with developmentMode as true
	// instance, err := vwo.Launch(settingsFile, api.WithDevelopmentMode(), api.WithStorage(storage), api.WithLogger(logs), api.WithGoalAttributes(nil, nil))

	settingsFile := vwo.GetSettingsFile(config.GetString("accountID"), config.GetString("SDKKey"))
	instance, err := vwo.Launch(settingsFile, api.WithDevelopmentMode(), api.WithStorage(storage), api.WithGoalAttributes(nil, nil))
	if err != nil {
		fmt.Println("error intialising sdk")
	}
	vwo.SetLogLevel(3)
	return instance
}
