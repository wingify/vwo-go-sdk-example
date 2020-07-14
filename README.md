# VWO Go SDK Example

This repository provides a basic demo of how server-side works with VWO GO SDK.

## Requirements

- Works with Go 1.1x

## Documentation

Refer [VWO Official FullStack Documentation](https://developers.vwo.com/reference#fullstack-introduction)

## Setup

1. Install dependencies

```go
go get .
go get -u github.com/cosmtrek/air
```

2. Update your app with your settings present in `config/dev.yaml`

```yaml
accountID: "REPLACE_THIS_WITH_CORRECT_VALUE"
SDKKey: "REPLACE_THIS_WITH_CORRECT_VALUE"
isDevelopmentMode: bool
abCampaignKey: "REPLACE_THIS_WITH_CORRECT_VALUE"
abCampaignGoalIdentifier: "REPLACE_THIS_WITH_CORRECT_VALUE"
featureRolloutCampaignKey: "REPLACE_THIS_WITH_CORRECT_VALUE"
featureTestCampaignKey: "REPLACE_THIS_WITH_CORRECT_VALUE"

revenueValue: "REPLACE_THIS_WITH_CORRECT_VALUE"

stringVariable: "REPLACE_THIS_WITH_CORRECT_VALUE"
intVariable: "REPLACE_THIS_WITH_CORRECT_VALUE"
boolVariable: "REPLACE_THIS_WITH_CORRECT_VALUE"
doubleVariable: "REPLACE_THIS_WITH_CORRECT_VALUE"
```

3. Run application

```shell
air
```

## Basic Usage

**Importing and Instantiation**

```go
import vwo "github.com/wingify/vwo-go-sdk"
import "github.com/wingify/vwo-go-sdk/pkg/api"

// Get SettingsFile
settingsFile := vwo.GetSettingsFile("accountID", "SDKKey")

// Default instance of vwoInstance
instance, err := vwo.Launch(settingsFile)
if err != nil {
	//handle err
}

// Instance with custom options
instance, err := vwo.Launch(settingsFile, api.WithDevelopmentMode())
if err != nil {
	//handle err
}

// Activate API
// With Custom Variables
options := make(map[string]interface{})
options["customVariables"] = map[string]interface{}{"a": "x"}
options["variationTargetingVariables"] = map[string]interface{}{}
options["revenueValue"] = 12
variationName = vwoInstance.Activate(campaignKey, userID, options)

// Without Custom Variables
variationName = instance.Activate(campaignKey, userID, nil)

// GetVariation
variationName = instance.GetVariationName(campaignKey, userID, nil)

// Track API with revenueValue in options
options := make(map[string]interface{})
options["revenueValue"] = 12
isSuccessful = instance.Track(campaignKey, userID, goalIdentifier, options)

// Track API with goalTypeToTrack and shouldTrackReturningUser in options
options := make(map[string]interface{})
options["goalTypeToTrack"] = constants.GoalType.All
options["shouldTrackReturningUser"] = false
isSuccessful = instance.Track(campaignKey, userID, goalIdentifier, options)

// FeatureEnabled API
isSuccessful = instance.IsFeatureEnabled(campaignKey, userID, nil)


// GetFeatureVariableValue API
variableValue = instance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)

// Push API
isSuccessful = instance.Push(tagKey, tagValue, userID)
```

**User Storage**

```go
import vwo "github.com/wingify/vwo-go-sdk/"
import "github.com/wingify/vwo-go-sdk/pkg/api"
import "github.com/wingify/vwo-go-sdk/pkg/schema"

// declare UserStorage interface with the following Get & Set function signature
type UserStorage interface{
    Get(userID, campaignKey string) UserData
    Set(string, string, string, string)
}

// declare a UserStorageData struct to implement UserStorage interface
type UserStorageData struct{}

// Get method to fetch user variation from storage
func (us *UserStorageData) Get(userID, campaignKey string) schema.UserData {
    //Example code showing how to get userData  from DB
    userData, ok := userDatas[campaignKey]
    if ok {
		for _, userdata := range userData {
			if userdata.UserID == userID {
				return userdata
			}
		}
    }
    /*
    // UserData  struct
    type UserData struct {
        UserID        string
        CampaignKey   string
        VariationName string
    }
    */
	return schema.UserData{}
}

// Set method to save user variation to storage
func (us *UserStorageData) Set(userID, campaignKey, variationName, goalIdentifer string) {
    //Example code showing how to store userData in DB
    userdata := schema.UserData{
		UserID:        userID,
		CampaignKey:   campaignKey,
		VariationName: variationName,
		GoalIdentifier: goalIdentifier,
	}

	flag := false
	userData, ok := userDatas[userdata.CampaignKey]
	if ok {
		for _, user := range userData {
			if user.UserID == userdata.UserID {
				flag = true
			}
		}
		if !flag {
			userDatas[userdata.CampaignKey] = append(userDatas[userdata.CampaignKey], userdata)
		} else {
			for i, user := range userData {
				if user.UserID == userdata.UserID && user.CampaignKey == userdata.CampaignKey {
					userData[i].VariationName = userdata.VariationName
					userData[i].GoalIdentifier = userdata.GoalIdentifier
				}
			}
		}
	} else {
		userDatas[userdata.CampaignKey] = []schema.UserData{
			userdata,
		}
	}
}

func main() {
	settingsFile := vwo.GetSettingsFile("accountID", "SDKKey")
	// create UserStorageData object
	storage := &UserStorageData{}

	instance, err := vwo.Launch(settingsFile, api.WithStorage(storage))
	if err != nil {
		//handle err
	}
}

```

**Custom Logger**

```go
import vwo "github.com/wingify/vwo-go-sdk"
import "github.com/wingify/vwo-go-sdk/pkg/api"

// declare Log interface with the following CustomLog function signature
type Log interface {
	CustomLog(level, errorMessage string)
}

// declare a LogS struct to implement Log interface
type LogS struct{}

// Get function to handle logs
func (c *LogS) CustomLog(level, errorMessage string) {}

func main() {
	settingsFile := vwo.GetSettingsFile("accountID", "SDKKey")
	// create LogS object
	logger := &LogS{}

	instance, err := vwo.Launch(settingsFile, api.WithLogger(logger))
	if err != nil {
		//handle err
	}
}
```

**Set log level**

```go
import vwo "github.com/wingify/vwo-go-sdk/"

vwo.SetLogLevel(3)
```

Levels

[0] Errors
[1] Warnings and Errors
[2] Info, Warnings and Errors
[3] Debugs, Info, Warnings and Errors

## Third-party Resources and Credits

Refer [third-party-attributions.txt](third-party-attributions.txt)

## License

[Apache License, Version 2.0](LICENSE)
