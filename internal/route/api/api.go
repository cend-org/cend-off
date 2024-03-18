package api

import (
	"duval/internal/pkg/address"
	"duval/internal/pkg/media"
	"duval/internal/pkg/user"
	"duval/internal/route/docs"
	"net/http"
)

var Routes = []docs.RouteDocumentation{
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/upload",
		Description:  "Uploads a media file to the server",
		IsPublic:     true,
		Handler:      media.Upload,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodHead,
		RelativePath: "/public",
		Description:  "",
		IsPublic:     true,
		DocRoot:      "public",
		NeedToken:    false,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/register/:as",
		Handler:      user.Register,
		IsPublic:     true,
		Description:  "Registers a new user with the specified user type (student , parent, tutor and professor)",
		NeedToken:    false,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/code/send",
		Handler:      user.SendUserEmailValidationCode,
		IsPublic:     true,
		Description:  "Sends a user email validation code for account validation",
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/code/verification/:code",
		Handler:      user.VerifyUserEmailValidationCode,
		IsPublic:     true,
		Description:  "Verifies a user email using the provided validation code",
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/login",
		Handler:      user.Login,
		IsPublic:     true,
		Description:  "Logs in a user with email and return token if success",
		NeedToken:    false,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/password",
		Handler:      user.NewPassword,
		IsPublic:     true,
		Description:  "Allows user to create password",
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/profile",
		Handler:      user.MyProfile,
		IsPublic:     true,
		Description:  "Retrieves information about the currently logged-in user's profile",
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPut,
		RelativePath: "/profile",
		Handler:      user.UpdMyProfile,
		IsPublic:     true,
		Description:  "Updates the currently logged-in user's profile information",
		NeedToken:    true,
	},
	// Address route
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/address",
		Handler:      address.NewAddress,
		IsPublic:     true,
		Description:  "Set the currently logged-in user's address",
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPut,
		RelativePath: "/address",
		Handler:      address.UpdateUserAddress,
		IsPublic:     true,
		Description:  "Edit the currently logged-in user's address",
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/address",
		Handler:      address.GetUserAddress,
		IsPublic:     true,
		Description:  "Get the currently logged-in user's address",
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/address",
		Handler:      address.RemoveUserAddress,
		IsPublic:     true,
		Description:  "Remove the currently logged-in user's address",
		NeedToken:    true,
	},
}
