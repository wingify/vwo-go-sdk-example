package models

//uncomment the commented imports to use UserStorage
import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "path/filepath"

	"github.com/wingify/vwo-go-sdk/pkg/schema"
)

// UserStorage interface
type UserStorage interface {
	Get(userID, campaignKey string) schema.UserData
	Set(userID, campaignKey, variationName, goalIdentifier string)
}

// UserStorageData struct
type UserStorageData struct{}

//Create a UserStorage.json file in models directory to use it for Get and Set
// path of example JSON user storage
var path = "./models/UserStorage.json"

// Get function is used to get the data from user storage
func (us *UserStorageData) Get(userID, campaignKey string) schema.UserData {
	var userDatas map[string][]schema.UserData

	// Conect your database here to fetch the current data
	// Uncomment the below part (Lines 33-45) to user JSON as data base

	// absPath, err := filepath.Abs(path)
	// if err != nil {
	// 	fmt.Println("Could not make absolute path: ", err)
	// }

	// file, err := ioutil.ReadFile(absPath)
	// if err != nil {
	// 	fmt.Println("Could not read userStorage: ", err)
	// }

	// if err = json.Unmarshal(file, &userDatas); err != nil {
	// 	fmt.Println("Could not unmarshall: ", err)
	// }

	if len(userDatas) == 0 {
		return schema.UserData{}
	}
	userData, ok := userDatas[campaignKey]
	if ok {
		for _, userdata := range userData {
			if userdata.UserID == userID {
				return userdata
			}
		}
	}
	return schema.UserData{}
}

// Set function
func (us *UserStorageData) Set(userID, campaignKey, variationName, goalIdentifier string) {
	var userDatas map[string][]schema.UserData

	// Conect your database here to insert the value
	// Uncomment the below part (Lines 68-80) to user JSON as data base

	// absPath, err := filepath.Abs(path)
	// if err != nil {
	// 	fmt.Println("Could not make absolute path: ", err)
	// }

	// file, err := ioutil.ReadFile(absPath)
	// if err != nil {
	// 	fmt.Println("Could not read userStorage: ", err)
	// }

	// if err = json.Unmarshal(file, &userDatas); err != nil {
	// 	fmt.Println("Could not unmarshall: ", err)
	// }

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

	// This is a part of the above JSON data handling

	// data, err := json.MarshalIndent(userDatas, "", " ")
	// if err != nil {
	// 	fmt.Println("Could not marshall: ", err)
	// }
	// err = ioutil.WriteFile(absPath, data, 0744)
	// if err != nil {
	// 	fmt.Println("Could not write userStorage: ", err)
	// }
}
