package main

import "github.com/BurntSushi/toml"

//Configuration . . . config contains useful information about configuration that can be used throughout the microservice.
type Configuration struct {
	Title        string
	Microservice microservice
	Routing      routing
}

type microservice struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
}

type routing struct {
	Server              string `toml:"ip"`
	Port                string `toml:"port"`
	HealthCheckStatusUp string `toml:"healthCheckStatusUp"`
}

//Config . . . exported config var holding configuration data
var Config Configuration

func init() {
	if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
		panic(err)
	}
}
