package main

import (
	"flag"
	app "myapp/app/server"
)

const (
	configFileKey     = "configFile"
	defaultConfigFile = ""
	configFileUsage   = "this is config file path"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, configFileKey, defaultConfigFile, configFileUsage)
	flag.Parse()

	app.Init(configFile)

	// for local run
	// app.Init("localDevSetup/local.yaml")
}
