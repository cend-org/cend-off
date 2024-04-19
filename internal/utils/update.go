package utils

import "github.com/mitchellh/mapstructure"

func ShouldBindJSON(changes map[string]interface{}, to interface{}) error {
	return mapstructure.Decode(changes, to)
}
