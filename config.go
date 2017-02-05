package main

import (
	"os"
	"encoding/json"
)

func GetConfig() Config {
	f, _ := os.Open("config.json")
	decoder := json.NewDecoder(f)
	config_obj := Config{}
	err := decoder.Decode(&config_obj)
	checkError(err)
	return config_obj
}