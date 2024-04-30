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
}

var App Config

func init() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "dev":
			App = Config{
				Version:              "",
				Port:                 "8087",
				Host:                 "",
				TokenSecret:          "a new token secret",
				DatabaseUserName:     "root",
				DatabaseUserPassword: "UnderAll4",
				DatabaseName:         "moja",
				DatabaseHost:         "localhost",
				DatabasePort:         "3306",
			}
		default:
			err := fullFillAppFromConfig()
			if err != nil {
				panic(err)
			}
		}
	} else {
		err := fullFillAppFromConfig()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(os.Args[1])

	App.DatabaseConnexionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", App.DatabaseUserName,
		App.DatabaseUserPassword, App.DatabaseHost, App.DatabasePort,
		App.DatabaseName)
}

func fullFillAppFromConfig() (err error) {
	_, err = toml.DecodeFile("./config.toml", &App)
	if err != nil {
		panic(err)
	}
	return err
}
