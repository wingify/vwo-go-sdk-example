package util

import (
	"math/rand"
)

// GetRandomUser function returns a random user from the array of pre defined users
func GetRandomUser() string {
	var users = []string{
		"Ashley",
		"Bill",
		"Chris",
		"Dominic",
		"Emma",
		"Faizan",
		"Gimmy",
		"Harry",
		"Ian",
		"John",
		"King",
		"Lisa",
		"Mona",
		"Nina",
		"Olivia",
		"Pete",
		"Queen",
		"Robert",
		"Sarah",
		"Tierra",
		"Una",
		"Varun",
		"Will",
		"Xin",
		"You",
		"Zeba",
	}

	ind := rand.Intn(24)
	return users[ind]
}
