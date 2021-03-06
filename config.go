package main

import (
	"encoding/json"
	"os"
)

// GetConfig read the config file to the Config object.
func GetConfig() Config {
	f, err := os.Open("./config.json")
	checkError(err)
	decoder := json.NewDecoder(f)
	configObj := Config{}
	err = decoder.Decode(&configObj)
	checkError(err)
	return configObj
}
