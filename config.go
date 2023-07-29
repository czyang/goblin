package main

import (
	"encoding/json"
	"os"
)

// GetConfig read the config file to the Config object.
func GetConfig(configPath string) Config {
	f, err := os.Open(configPath)
	checkError(err)
	decoder := json.NewDecoder(f)
	configObj := Config{}
	err = decoder.Decode(&configObj)
	checkError(err)
	return configObj
}
