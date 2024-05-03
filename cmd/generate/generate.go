//go:generate go run generate.go

package main

import (
	"fmt"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/cend-org/duval/cmd/generate/mutation"
	"os"
)

func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	b = mutation.MutationHook(b)
	//b = migrate.MigrationHook(b)
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
