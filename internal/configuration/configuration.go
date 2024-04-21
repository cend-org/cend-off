package configuration

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

var App Configuration

type Configuration struct {
	Version                 string   `yaml:"version"`
	RunningMode             string   `yaml:"running_mode"`
	Database                Database `yaml:"database"`
	RunHash                 string   `yaml:"run_hash"`
	DatabaseConnexionString string
}

type Database struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	UserName     string `yaml:"user_name"`
	UserPassword string `yaml:"user_password"`
	Name         string `yaml:"name"`
}

func init() {
	App = local

	App.DatabaseConnexionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", App.Database.UserName,
		App.Database.UserPassword, App.Database.Host, App.Database.Port,
		App.Database.Name)

	App.RunHash = uuid.NewV4().String()

}

func (app *Configuration) IsDebug() bool {
	return app.RunningMode == "DEBUG"
}

func (app *Configuration) IsProd() bool {
	return app.RunHash == "PROD"
}
