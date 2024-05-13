package main

import (
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
	migrate "github.com/cend-org/duval/cmd/generate/migration"
	"github.com/cend-org/duval/cmd/generate/mutation"
	"github.com/iancoleman/strcase"
)

type iMutator struct {
	hooks    []func(b *modelgen.ModelBuild) *modelgen.ModelBuild
	mutate   func(b *modelgen.ModelBuild) *modelgen.ModelBuild
	generate func()
}

var dummy = iMutator{}

func init() {
	dummy.hooks = []func(b *modelgen.ModelBuild) *modelgen.ModelBuild{
		func(b *modelgen.ModelBuild) *modelgen.ModelBuild {
			for _, model := range b.Models {
				for _, field := range model.Fields {
					snakeCaseName := strcase.ToLowerCamel(field.Name)
					field.Tag = `json:"` + snakeCaseName + `"`
				}
			}
			return b
		},
		mutation.Hook,
		migrate.MigrationHook,
	}

	dummy.mutate = func(b *modelgen.ModelBuild) *modelgen.ModelBuild {
		for i := 0; i < len(dummy.hooks); i++ {
			b = dummy.hooks[i](b)
		}
		return b
	}

	dummy.generate = func() {
		cfg, err := config.LoadConfigFromDefaultLocations()
		if err != nil {
			panic(err)
		}

		p := modelgen.Plugin{
			MutateHook: dummy.mutate,
		}

		err = api.Generate(cfg,
			api.ReplacePlugin(&p),
		)

		if err != nil {
			panic(err)
		}
	}
}

func main() { dummy.generate() }
