//go:build ignore

package main

import (
	"fmt"
	"os"
	"unicode"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

func camelToSnake(camelCase string) string {
	var result []rune

	for i, char := range camelCase {
		if unicode.IsUpper(char) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(char))
		} else {
			result = append(result, char)
		}
	}

	return string(result)
}

func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			snakeCaseName := camelToSnake(field.Name)
			field.Tag = `json:"` + snakeCaseName + `"`
		}
	}

	return b
}

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	err = api.Generate(cfg, api.ReplacePlugin(&p))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
