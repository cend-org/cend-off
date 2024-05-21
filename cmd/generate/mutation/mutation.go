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
	for _, model := range b.Models {
		if !strings.Contains(model.Name, "Input") {
			continue
		}

		createFieldValidation(model)
		var raw string
		raw += createObjectEntityCollector(model)
		raw += createInputMap(model)

		fmt.Println("- generating ", model.Name, " mapping func ...")

		// create file and write raw content
		f, err := os.Create(fmt.Sprintf("./graph/model/%s_gen.go", strcase.ToSnake(model.Name)))
		if err != nil {
			log.Fatal(err)
		}

		defer func(f *os.File) {
			errF := f.Close()
			if errF != nil {
				panic(errF)
			}
		}(f)

		_, err = f.WriteString(raw)
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

func createInputMap(model *modelgen.Object) (input string) {
	modelTypeName := strings.ReplaceAll(model.Name, "Input", "")
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

func createObjectEntityCollector(model *modelgen.Object) (raw string) {
	modelTypeName := strings.ReplaceAll(model.Name, "Input", "")
	raw += fmt.Sprintf("package model\n\n")
	raw += fmt.Sprint("import (\n")
	raw += fmt.Sprintf("		\"github.com/cend-org/duval/internal/database\"\n")
	raw += fmt.Sprintf("		\"time\"\n")
	raw += fmt.Sprint(")\n\n")

	raw += fmt.Sprintf("type %sCollector struct {} \n", strcase.ToCamel(modelTypeName))
	for _, field := range model.Fields {
		if strings.Contains(field.Tag, "gql") {
			continue
		}
		raw += fmt.Sprintf("func (c *%sCollector) %s(a %s) (r %s, err error) { \n", strcase.ToCamel(modelTypeName), strcase.ToCamel(field.Name), field.Type.String(), modelTypeName)
		raw += fmt.Sprintf("		err = database.Get(&r, `SELECT * FROM %s WHERE %s = ? ORDER BY created_at DESC LIMIT 1`, a)\n", strcase.ToSnake(modelTypeName), strcase.ToSnake(field.Name))
		raw += fmt.Sprintf("  return r, err\n")
		raw += fmt.Sprintf("}\n\n")
	}

	raw += fmt.Sprintf("func (c *%sCollector) _b(){ \n", strcase.ToCamel(modelTypeName))
	raw += fmt.Sprintf("		_ = time.Now()\n")
	raw += fmt.Sprintf("}\n\n")

	return raw
}

func createFieldValidation(model *modelgen.Object) {
	for _, field := range model.Fields {
		switch strings.ToLower(field.Name) {
		case "email":
			field.Tag += ` validate:"email"`
		case "password":
			field.Tag += ` validate:"min=4"`
		case "age":
			field.Tag += ` validate:"gte=18,lte=130"`
		default:
			if strings.Contains(strings.ToLower(field.Name), "name") {
				field.Tag += `validate:"alpha"`
			} else {
				field.Tag += ``
			}
		}
	}
}
