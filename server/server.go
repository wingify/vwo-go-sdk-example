package server

import "github.com/wingify/vwo-go-sdk/config"

//Init function starts the server on the port specified
func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("port"))
}
