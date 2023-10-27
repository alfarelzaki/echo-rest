package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func GetConfig() Configuration {
	// create an empty Configuration struct
	conf := Configuration{}

	// trying to get configuration data from config.json
	gonfig.GetConf("config/config.json", &conf)
	
	return conf
}
