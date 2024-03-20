package api

import (
	"duval/internal/pkg/translator"
	"duval/internal/route/docs"
	"net/http"
)

var TranslationRoutes = []docs.RouteDocumentation{
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/messages",
		Handler:      translator.GetMessages,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/messages/:language",
		NeedToken:    false,
		Handler:      translator.GetMessagesInLanguage,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/message/:number/:language",
		NeedToken:    false,
		Handler:      translator.GetMessage,
		DocRoot:      "",
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/message",
		Handler:      translator.NewMessage,
		NeedToken:    false,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/message/:number/:language",
		Handler:      translator.DelMessage,
		NeedToken:    false,
	},
	{
		HttpMethod:   http.MethodPut,
		RelativePath: "/message",
		NeedToken:    false,
		Handler:      translator.UpdMessage,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/menu",
		NeedToken:    false,
		Handler:      translator.GetMenuList,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/menu/item/:number/:language",
		NeedToken:    false,
		Handler:      translator.GetMenuItems,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/menu/:language",
		NeedToken:    false,
		Handler:      translator.NewMenu,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/menu/item",
		NeedToken:    false,
		Handler:      translator.NewMenuItem,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/menu/:number",
		NeedToken:    false,
		Handler:      translator.DelMenu,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/menu/item",
		NeedToken:    false,
		Handler:      translator.DelMenuItem,
	},
}
