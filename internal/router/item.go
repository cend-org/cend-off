package router

import (
	"github.com/cend-org/duval/internal/router/api"
	"github.com/cend-org/duval/internal/router/docs"
)

var RootRoutesGroup = []docs.RootDocumentation{
	{
		Group: "",
		Paths: api.Routes,
	},
}
