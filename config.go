package main

import (
	"encoding/json"
	"os"
)

// GetConfig reads the config file into a Config object.
func GetConfig(configPath string) Config {
	f, err := os.Open(configPath)
	if err != nil {
		logFatalIfError(err)
	}
	defer f.Close()

	configObj := Config{}
	err = json.NewDecoder(f).Decode(&configObj)
	if err != nil {
		logFatalIfError(err)
	}

	return configObj
}
