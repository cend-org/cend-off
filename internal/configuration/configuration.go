package configuration

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

const (
	DEV_MODE  = "dev"
	TEST_MODE = "test"
	PROD_MODE = "prod"
)

type Config struct {
	Version                 string `toml:"version"`
	Port                    string `toml:"port"`
	Host                    string `toml:"host"`
	TokenSecret             string `toml:"token_secret"`
	DatabaseUserName        string `toml:"database_user_name"`
	DatabaseUserPassword    string `toml:"database_user_password"`
	DatabaseName            string `toml:"database_name"`
	DatabaseHost            string `toml:"database_host"`
	DatabasePort            string `toml:"database_port"`
	DatabaseConnexionString string
	Mode                    string `toml:"mode"`
}

var App Config

func init() {
	err := fullFillAppFromConfig()
	if err != nil {
		panic(err)
	}

	App.DatabaseConnexionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", App.DatabaseUserName,
		App.DatabaseUserPassword, App.DatabaseHost, App.DatabasePort,
		App.DatabaseName)

	/*
		APP.MODE is set to test mode by default
	*/
	if len(App.Mode) == 0 {
		App.Mode = TEST_MODE
	}
}

func fullFillAppFromConfig() (err error) {
	_, err = toml.DecodeFile("./config.toml", &App)
	if err != nil {
		panic(err)
	}

	fmt.Sprintln("App is starting with :")
	fmt.Sprintln(App)

	return err
}

func IsDev() bool {
	return App.Mode == DEV_MODE
}

func IsProd() bool {
	return App.Mode == PROD_MODE
}

func IsTest() bool {
	return App.Mode == TEST_MODE
}
