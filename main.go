package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wingify/vwo-go-example/config"
	"github.com/wingify/vwo-go-example/server"
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
