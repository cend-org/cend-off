package mutation

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/iancoleman/strcase"
)

func Hook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	var (
		inputPath string
	)

	for _, model := range b.Models {
		if !strings.Contains(model.Name, "Input") {
			continue
		}

		inputPath = fmt.Sprintf("./graph/model/%s_gen.go", strcase.ToSnake(model.Name))

		f, err := os.Create(inputPath)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		s := inputString(model)

		_, err = f.WriteString(s)
		if err != nil {
			log.Fatal(err)
		}

		err = f.Sync()
		if err != nil {
			panic(err)
		}
	}

	return b
}

func inputString(model *modelgen.Object) (input string) {
	fmt.Println("- generating ", model.Name, " mapping func ...")

	input += fmt.Sprintf("package model\n")
	modelTypeName := strings.ReplaceAll(model.Name, "Input", "")
	input += fmt.Sprintf("\n")
	input += fmt.Sprintf("func Map%sTo%s(a %s, e %s) %s { \n", model.Name, modelTypeName, model.Name, modelTypeName, modelTypeName)
	for _, field := range model.Fields {
		if strings.Contains(field.Tag, "gql") {
			continue
		}
		input += fmt.Sprintf(" if a.%s != nil { \n", field.Name)
		input += fmt.Sprintf("	e.%s = *a.%s \n", field.Name, field.Name)
		input += fmt.Sprintf(" }\n")
	}
	input += fmt.Sprintf("  return e\n")
	input += fmt.Sprintf("}\n")
	return input
}
