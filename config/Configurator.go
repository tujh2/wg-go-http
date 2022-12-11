package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var configPath = "./"
var configName = "wg-go-http.conf"

type Config struct {
	HttpHost string
	Secret   string
}

func ReadConfig() Config {
	configBytes, err := os.ReadFile(configPath + configName)
	if err != nil {
		log.Fatal("ERROR: Failed to read config file:\n" + err.Error())
	}
	tomlData := string(configBytes)
	var config Config
	if _, err := toml.Decode(tomlData, &config); err != nil {
		log.Fatal("ERROR: Failed to parse config file:\n " + err.Error())
	}
	return config
}
