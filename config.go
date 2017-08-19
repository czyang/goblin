package main

import (
	"encoding/json"
	"os"
)

func GetConfig() Config {
	f, _ := os.Open("config.json")
	decoder := json.NewDecoder(f)
	config_obj := Config{}
	err := decoder.Decode(&config_obj)
	checkError(err)
	return config_obj
}
