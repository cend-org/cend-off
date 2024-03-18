package api

import (
	"duval/internal/pkg/translator"
	"duval/internal/route/docs"
	"net/http"
)

var TranslationRoutes = []docs.RouteDocumentation{
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/message",
		Description:  "",
		IsPublic:     true,
		Handler:      translator.NewMessage,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPut,
		RelativePath: "/message",
		Description:  "",
		IsPublic:     true,
		Handler:      translator.UpdMessage,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/message",
		Description:  "",
		IsPublic:     true,
		Handler:      translator.DeleteMessage,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/messages",
		Description:  "",
		IsPublic:     true,
		Handler:      translator.Messages,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/message",
		Description:  "",
		IsPublic:     true,
		Handler:      translator.Message,
	},
}
