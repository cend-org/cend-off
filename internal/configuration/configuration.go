package configuration

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"strconv"
)

const (
	DEV_MODE  = "dev"
	TEST_MODE = "test"
	PROD_MODE = "prod"
)

type Config struct {
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

	verifyAppConfiguration()

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

func verifyAppConfiguration() {
	if value, err := strconv.Atoi(App.Port); len(App.Port) == 0 || err != nil || value == 0 {
		fmt.Println("❌ PORT value must be provided")
		os.Exit(1)
	} else {
		fmt.Println("✔️ PORT SET ON :", App.Port)
	}

	if len(App.TokenSecret) == 0 {
		fmt.Println("❌ TOKEN SECRET is not defined !")
		os.Exit(1)
	} else {
		fmt.Println("✔️ TOKEN SECRET  :", App.TokenSecret)
	}

	if len(App.DatabaseUserName) == 0 {
		fmt.Println("❌ DB USER NAME must be provided !")
		os.Exit(1)
	} else {
		fmt.Println("✔️ DB USER NAME  :", App.DatabaseUserName)
	}

	if len(App.DatabasePort) == 0 {
		fmt.Println("❌ DB PORT must be provided !")
		os.Exit(1)
	} else {
		fmt.Println("✔️ DB PORT :", App.DatabasePort)
	}

	if len(App.DatabaseUserPassword) == 0 {
		fmt.Println("❌ DB USER PASSWORD must be provided !")
		os.Exit(1)
	} else {
		fmt.Println("✔️ DB PASSWORD :", App.DatabaseUserPassword)
	}
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
