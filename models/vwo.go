package models

import (
	"fmt"
	vwo "github.com/wingify/vwo-go-sdk"
	"github.com/wingify/vwo-go-sdk/pkg/api"
	"log"
	"wingify.com/vwo-go-sdk-example/config"
)

// GetVWOInstance function initializes the settings file and launches the
// VWO instance using configuration values and returns the insttance

var instance *api.VWOInstance
var err error

func GetVWOInstance() *api.VWOInstance {
	if instance != nil {
		return instance
	}

	config := config.GetConfig()

	settingsFile := vwo.GetSettingsFile(config.GetString("accountID"), config.GetString("SDKKey"))

	// Event Batching
	callBack := func(err error, batch []map[string]interface{}) {
		log.Println(fmt.Sprintf("Batch events pushed %v %v", batch, err))
	}

	// Integrations callBack
	Integrations := func(integrationsMap map[string]interface{}) {
		fmt.Println("Integrations Map : ", integrationsMap)
	}

	// User Storage
	//storage := &UserStorageData{}

	// Custom Logger (Google Logger)
	// logs := &LogS{}

	// Instance with userStorage and customLogger with developmentMode as true
	// instance, err := vwo.Launch(settingsFile, api.WithDevelopmentMode(), api.WithStorage(storage), api.WithLogger(logs), api.WithGoalAttributes(nil, nil))

	instance, err = vwo.Launch(
		settingsFile
		// api.WithBatchEventQueue(api.BatchConfig{RequestTimeInterval: 10, EventsPerRequest: 10}, callBack),
		// api.WithIntegrationsCallBack(Integrations)
	)

	if err != nil {
		fmt.Println("error intialising sdk")
	}

	vwo.SetLogLevel(3)

	return instance
}
