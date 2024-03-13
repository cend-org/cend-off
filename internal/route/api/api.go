package api

import (
	"duval/internal/pkg/media"
	"duval/internal/pkg/user"
	"duval/internal/route/docs"
	"net/http"
)

var Routes = []docs.RouteDocumentation{
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/upload",
		Description:  "",
		IsPublic:     true,
		Handler:      media.Upload,
	},
	{
		HttpMethod:   http.MethodHead,
		RelativePath: "/public",
		Description:  "",
		IsPublic:     true,
		DocRoot:      "public",
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/register/:as",
		Handler:      user.Register,
		IsPublic:     true,
		Description:  "",
	},
}
