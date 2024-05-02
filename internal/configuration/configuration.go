package configuration

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
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
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case DEV_MODE:
			/*
				The dev mod is used for dev release. If You are using the app on localhost, please use the test mode.
			*/
			App = Config{
				Version:              "",
				Port:                 "8087",
				Host:                 "",
				TokenSecret:          "a new token secret",
				DatabaseUserName:     "root",
				DatabaseUserPassword: "UnderAll4",
				DatabaseName:         "cend",
				DatabaseHost:         "cend.ctw4aeiceahd.eu-north-1.rds.amazonaws.com",
				DatabasePort:         "3306",
				Mode:                 DEV_MODE,
			}
		default:
			err := fullFillAppFromConfig()
			if err != nil {
				panic(err)
			}
		}

		App.DatabaseConnexionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", App.DatabaseUserName,
			App.DatabaseUserPassword, App.DatabaseHost, App.DatabasePort, App.DatabaseName)
	} else {
		err := fullFillAppFromConfig()
		if err != nil {
			panic(err)
		}

		App.DatabaseConnexionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", App.DatabaseUserName,
			App.DatabaseUserPassword, App.DatabaseHost, App.DatabasePort,
			App.DatabaseName)
	}

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
