package validator

import (
	"fmt"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/iancoleman/strcase"
	"log"
	"os"
	"strings"
)

type FieldPattern struct {
	Identifier string
	Pattern    string
}

func ValidationHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		if !strings.Contains(model.Name, "Input") {
			continue
		}

		var raw string

		raw += createValidationServices(model)

		fmt.Println("- generating ", model.Name, " validation func ...")

		// create file and write raw content
		f, err := os.Create(fmt.Sprintf("./internal/utils/validators/%s_validator_gen.go", strcase.ToSnake(model.Name)))
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

func createValidationServices(model *modelgen.Object) (input string) {
	patterns := []FieldPattern{
		{Identifier: "email", Pattern: `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`},
		{Identifier: "password", Pattern: `^(?=.[A-Za-z])(?=.\d)[A-Za-z\d@$!%*?&]{4,}$`},
		{Identifier: "date", Pattern: `^\d{4}-\d{2}-\d{2}$`},
		{Identifier: "name", Pattern: `^[a-zA-Z]+$`},
		{Identifier: "phone", Pattern: `^\+?[1-9]\d{1,14}$`},
		{Identifier: "url", Pattern: `^(https?|ftp)://[^\s/$.?#].[^\s]*$`},
		{Identifier: "ipv4", Pattern: `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`},
		{Identifier: "ipv6", Pattern: `^([0-9a-fA-F]{1,4}:){7}([0-9a-fA-F]{1,4}|:)$`},
		{Identifier: "uuid", Pattern: `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`},
	}

	input += fmt.Sprintf("package utils \n \n")
	input += fmt.Sprintf("import (\n\t\"errors\"\n\t\"github.com/cend-org/duval/graph/model\"\n\t\"regexp\"\n\t\"time\"\n) \n")

	input += fmt.Sprintf("func %sValidator(i model.%s) (err error) { \n", model.Name, model.Name)
	for _, field := range model.Fields {
		if strings.Contains(field.Tag, "gql") {
			continue
		}
		modelTypeName := strings.ReplaceAll(model.Name, "Input", "")
		input += fmt.Sprintf("	if IsValid%s%s(i.%s) { \n", strcase.ToCamel(modelTypeName), strcase.ToCamel(field.Name), strcase.ToCamel(field.Name))
		input += fmt.Sprintf("		return errors.New(\"%s %s input invalid\") \n", strings.ToLower(modelTypeName), strcase.ToLowerCamel(field.Name))
		input += fmt.Sprintf("	} \n")
	}

	input += fmt.Sprintf("	return nil \n")
	input += fmt.Sprintf("} \n")

	for _, field := range model.Fields {
		if strings.Contains(field.Tag, "gql") {
			continue
		}

		var pattern string
		for _, fp := range patterns {
			if strings.Contains(strings.ToLower(field.Name), fp.Identifier) {
				pattern = fp.Pattern
				break
			}
		}

		// Default to alphanumeric validation if no specific pattern is matched
		if pattern == "" {
			pattern = `^[a-zA-Z0-9]+$`
		}

		modelTypeName := strings.ReplaceAll(model.Name, "Input", "")
		input += fmt.Sprintf("func IsValid%s%s(%s %s) bool {\n", strcase.ToCamel(modelTypeName), strcase.ToCamel(field.Name), strcase.ToLowerCamel(field.Name), field.Type)
		input += fmt.Sprintf("\tpattern := `%s`\n", pattern)
		input += fmt.Sprintf("\tregex := regexp.MustCompile(pattern)\n")
		input += fmt.Sprintf("\treturn regex.MatchString(*%s)\n", strcase.ToLowerCamel(field.Name))
		input += fmt.Sprintf("}\n\n")

	}

	return input
}
