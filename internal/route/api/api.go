package api

import (
	"duval/internal/pkg/address"
	"duval/internal/pkg/media"
	"duval/internal/pkg/media/cover"
	cvtype "duval/internal/pkg/media/cv"
	"duval/internal/pkg/media/profile"
	video "duval/internal/pkg/media/video"
	"duval/internal/pkg/user"
	"duval/internal/pkg/user/link"
	"duval/internal/route/docs"
	"net/http"
)

var Routes = []docs.RouteDocumentation{
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/upload",
		Handler:      media.Upload,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodHead,
		RelativePath: "/public",
		DocRoot:      "public",
		NeedToken:    false,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/register/:as",
		Handler:      user.Register,
		NeedToken:    false,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/code/send",
		Handler:      user.SendUserEmailValidationCode,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/code/verification/:code",
		Handler:      user.VerifyUserEmailValidationCode,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/login",
		Handler:      user.Login,
		NeedToken:    false,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/password",
		Handler:      user.NewPassword,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/profile",
		Handler:      user.MyProfile,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPut,
		RelativePath: "/profile",
		Handler:      user.UpdMyProfile,
		NeedToken:    true,
	},
	// Address route
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/address",
		Handler:      address.NewAddress,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPut,
		RelativePath: "/address",
		Handler:      address.UpdateUserAddress,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/address",
		Handler:      address.GetUserAddress,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/address",
		Handler:      address.RemoveUserAddress,
		NeedToken:    true,
	},
	//Profile image routes
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/profile/image",
		Handler:      profile.Upload,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPut,
		RelativePath: "/profile/image",
		Handler:      profile.UpdateProfileImage,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/profile/image",
		Handler:      profile.GetProfileImage,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/profile/thumb",
		Handler:      profile.GetProfileThumb,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/profile/image",
		Handler:      profile.RemoveProfileImage,
		NeedToken:    true,
	},
	//	link between user routes
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/user/parent",
		Handler:      link.GetUserParent,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/user/tutor",
		Handler:      link.GetUserTutor,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/user/professor",
		Handler:      link.GetUserProfessor,
		NeedToken:    true,
	},

	//Cover Presentation routes
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/profile/letter",
		Handler:      cover.UploadCoverLetter,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPut,
		RelativePath: "/profile/letter",
		Handler:      cover.UpdateProfileLetter,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/profile/letter",
		Handler:      cover.GetProfileLetter,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/profile/letter/thumb",
		Handler:      cover.GetProfileLetterThumb,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/profile/letter",
		Handler:      cover.RemoveProfileLetter,
		NeedToken:    true,
	},
	//cv_type  routes
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/profile/cv",
		Handler:      cvtype.UploadCv,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPut,
		RelativePath: "/profile/cv",
		Handler:      cvtype.UpdateProfileCv,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/profile/cv",
		Handler:      cvtype.GetProfileCv,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/profile/cv/thumb",
		Handler:      cvtype.GetProfileCvThumb,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/profile/cv",
		Handler:      cvtype.RemoveProfileCv,
		NeedToken:    true,
	},
	//videos  routes
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/profile/video",
		Handler:      video.UploadVideo,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPut,
		RelativePath: "/profile/video",
		Handler:      video.UpdateProfileVideo,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/profile/video",
		Handler:      video.GetProfileVideo,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/profile/video",
		Handler:      video.RemoveProfileVideo,
		NeedToken:    true,
	},
}
