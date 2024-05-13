package api

import (
	"github.com/cend-org/duval/internal/router/docs"
	mediafile "github.com/cend-org/duval/pkg/media"
	"net/http"
)

var Routes = []docs.RouteDocumentation{
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/upload",
		Handler:      mediafile.Upload,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodHead,
		RelativePath: "/public",
		DocRoot:      "public",
		NeedToken:    false,
	},
}
