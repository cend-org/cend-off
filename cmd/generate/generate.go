//go:generate go run generate.go

package main

import (
	"fmt"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"log"
	"os"
	"strings"
	"unicode"
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

// Defining mutation function
func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	inputPath := "./graph/model/models_input_gen.go"

	f, err := os.Create(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := inputString(b.Models)

	_, err = f.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Sync()
	if err != nil {
		panic(err)
	}

	for _, model := range b.Models {
		for _, field := range model.Fields {
			snakeCaseName := camelToSnake(field.Name)
			field.Tag = `json:"` + snakeCaseName + `"`
		}
	}

	return b
}

func main() {
	fmt.Println("Generating...")
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	err = api.Generate(cfg,
		api.ReplacePlugin(&p),
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

}

func inputString(s []*modelgen.Object) (input string) {

	input += fmt.Sprintf("package model \n")

	for _, model := range s {
		if !strings.Contains(model.Name, "Input") {
			continue
		}

		fmt.Println("- generating ", model.Name, " mapping func ...")

		input += fmt.Sprintf("\n\n /* %s */ \n\n", model.Name)

		modelTypeName := strings.ReplaceAll(model.Name, "Input", "")

		input += fmt.Sprintf("\n\n")

		input += fmt.Sprintf("func Map%sTo%s(input %s, existing %s) %s { \n", model.Name, modelTypeName, model.Name, modelTypeName, modelTypeName)

		for _, field := range model.Fields {
			if strings.Contains(field.Tag, "gql") {
				continue
			}

			input += fmt.Sprintf(" if input.%s != nil { \n", field.Name)
			input += fmt.Sprintf("	existing.%s = *input.%s \n", field.Name, field.Name)
			input += fmt.Sprintf(" } \n \n")

		}

		input += fmt.Sprintf("  return existing \n")

		input += fmt.Sprintf("}")
	}

	return input
}
