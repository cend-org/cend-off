package migrate

import (
	"fmt"
	"log"
	"os"
	"strings"

	"reflect"

	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/iancoleman/strcase"
)

func MigrationHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	inputPath := "./migration/db/migration_generated.sql"

	f, err := os.Create(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := migrationString(b.Models)

	_, err = f.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Sync()
	if err != nil {
		panic(err)
	}

	return b
}

func migrationString(b []*modelgen.Object) (migration string) {
	migration += "drop database if exists cend;\n"
	migration += "create database cend;\n"
	migration += "use cend;\n\n"

	for _, model := range b {
		if strings.HasSuffix(strcase.ToSnake(model.Name), "_input") {
			continue
		}
		migration += fmt.Sprintf("create table %s (\n", strcase.ToSnake(model.Name))

		migration += "\tid int auto_increment primary key,\n"

		for _, field := range model.Fields {
			if strcase.ToSnake(field.Name) == "id" {
				continue
			}
			migration += fmt.Sprintf("\t%s", strcase.ToSnake(field.Name))
			migration += "\t"
			migration += determineDataType(field.Type.String())
			migration += " "
			if strcase.ToSnake(field.Name) == "created_at" {
				migration += fmt.Sprintf("default %s", setDefaultTimeValue())

			} else if strcase.ToSnake(field.Name) == "updated_at" {
				migration += fmt.Sprintf("default %s", setDefaultTimeValue())

			} else {
				migration += fmt.Sprintf("default %s", setDefaultValue(field.Type.String()))
			}

			tag := reflect.StructTag(field.Tag)
			dbTag, ok := tag.Lookup("db")
			if ok {
				migration += " " + dbTag
			}

			migration += ",\n"
		}

		if len(model.Fields) > 0 {
			migration += fmt.Sprintf("\tconstraint %s_pk\n", strcase.ToSnake(model.Name))
			migration += fmt.Sprintf("\t\tunique (%s)\n", strcase.ToSnake(model.Fields[0].Name))
		}

		migration += ");\n\n"

		for _, field := range model.Fields {
			if strings.HasSuffix(strcase.ToSnake(field.Name), "_id") {
				referenceTable := setRefTable(strings.TrimSuffix(strcase.ToSnake(field.Name), "_id"))
				migration += fmt.Sprintf("alter table %s\n", strcase.ToSnake(model.Name))
				migration += fmt.Sprintf("\tadd constraint %s_%s_fk\n", strcase.ToSnake(model.Name), strcase.ToSnake(field.Name))
				migration += fmt.Sprintf("\t\tforeign key (%s) references %s (id);\n\n", strcase.ToSnake(field.Name), referenceTable)
			}

			if strings.HasSuffix(strcase.ToSnake(field.Name), "_xid") {
				referenceTable := setRefTable(strings.TrimSuffix(strcase.ToSnake(field.Name), "_xid"))
				migration += fmt.Sprintf("alter table %s\n", strcase.ToSnake(model.Name))
				migration += fmt.Sprintf("\tadd constraint %s_%s_fk\n", strcase.ToSnake(model.Name), strcase.ToSnake(field.Name))
				migration += fmt.Sprintf("\t\tforeign key (%s) references %s (xid);\n\n", strcase.ToSnake(field.Name), referenceTable)
			}
		}

	}

	return migration
}

func determineDataType(fieldType string) string {
	switch fieldType {
	case "time.Time", "*time.Time":
		return "datetime"
	case "string", "*string":
		return "varchar(255)"
	case "float64", "*float64":
		return "float"
	case "int", "*int":
		return "int"
	case "Date":
		return "date"
	default:
		return strcase.ToSnake(fieldType)
	}
}

func setDefaultValue(fieldType string) string {
	switch fieldType {
	case "time.Time", "*time.Time":
		return `'0000-00-00 00:00:00'`
	case "string", "*string":
		return `''`
	case "float64", "*float64":
		return "0"
	case "int", "*int":
		return "0"
	case "Date":
		return `'0000-00-00'`
	default:
		return `''`
	}
}

func setDefaultTimeValue() string {
	return "CURRENT_TIMESTAMP"
}

func setRefTable(refTable string) string {
	switch refTable {
	case "owner", "user", "publisher", "tutor", "parent", "student", "professor", "viewer", "author":
		return "user"
	default:
		return refTable
	}
}
