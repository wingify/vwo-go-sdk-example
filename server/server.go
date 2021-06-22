package server

import "wingify.com/vwo-go-sdk-example/config"

//Init function starts the server on the port specified
func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("port"))
}
