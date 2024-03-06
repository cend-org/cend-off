package configuration

import "github.com/BurntSushi/toml"

type Config struct {
	Version     string `toml:"version"`
	Port        string `toml:"port"`
	Host        string `toml:"host"`
	TokenSecret string `toml:"token_secret"`
	Database    string `toml:"database"`
}

var App Config

func init() {
	_, err := toml.DecodeFile("./config.toml", &App)
	if err != nil {
		panic(err)
	}
}
