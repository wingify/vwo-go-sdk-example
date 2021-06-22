package main

import (
	"flag"
	"fmt"
	"os"

	"wingify.com/vwo-go-sdk-example/config"
	"wingify.com/vwo-go-sdk-example/server"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}
