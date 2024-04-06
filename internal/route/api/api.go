package api

import (
	"duval/internal/authentication"
	"duval/internal/pkg/address"
	"duval/internal/pkg/education"
	"duval/internal/pkg/media"
	"duval/internal/pkg/media/cover"
	cvtype "duval/internal/pkg/media/cv"
	"duval/internal/pkg/media/profile"
	video "duval/internal/pkg/media/video"
	"duval/internal/pkg/planning"
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
		RelativePath: "/register/:as/:email",
		Handler:      user.RegisterByEmail,
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
		HttpMethod:   http.MethodGet,
		RelativePath: "/code",
		Handler:      user.GetCode,
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
		RelativePath: "/password/history",
		Handler:      user.GetPasswordHistory,
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
		RelativePath: "/profile/active",
		Handler:      user.ActivateUser,
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
		HttpMethod:   http.MethodPost,
		RelativePath: "/user/parent",
		Handler:      link.AddParentToUser,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/user/parent",
		Handler:      link.GetUserParent,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/user/parent",
		Handler:      link.RemoveUserParent,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/user/tutor",
		Handler:      link.AddTutorToUser,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/user/tutor",
		Handler:      link.GetUserTutor,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/user/tutor",
		Handler:      link.RemoveUserTutor,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/user/professor",
		Handler:      link.AddProfessorToUser,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/user/professor",
		Handler:      link.GetUserProfessor,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/user/professor",
		Handler:      link.RemoveUserProfessor,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/user/student",
		Handler:      link.AddStudentToLink,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/user/student",
		Handler:      link.GetStudent,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/user/student",
		Handler:      link.RemoveStudent,
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
	//Qr code authentication
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/generate-qr",
		Handler:      authentication.GenerateQrCode,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPut,
		RelativePath: "/login/with-qr/:xid",
		Handler:      authentication.LoginWithQr,
		NeedToken:    false,
	},
	//Calendar planning routes
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/calendar",
		Handler:      planning.CreateUserPlannings,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/calendar",
		Handler:      planning.GetUserPlannings,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/calendar",
		Handler:      planning.RemoveUserPlannings,
		NeedToken:    true,
	},
	//Calendar user planning routes
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/calendar/:calendar_id/:actor",
		Handler:      planning.AddUserIntoPlanning,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/calendar/:calendar_id/actor",
		Handler:      planning.GetPlanningActors,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodDelete,
		RelativePath: "/calendar/:calendar_id/actor",
		Handler:      planning.RemoveUserFromPlanning,
		NeedToken:    true,
	},
	//Education Routes
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/education",
		Handler:      education.GetEducation,
		NeedToken:    false,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/education/:edu",
		Handler:      education.GetSubjects,
		NeedToken:    false,
	},
	//Education Level Routes'
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/user/education/",
		Handler:      education.SetUserEducationLevel,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/user/education",
		Handler:      education.GetUserEducationLevel,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodPut,
		RelativePath: "/user/education/",
		Handler:      education.UpdateUserEducationLevel,
		NeedToken:    true,
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/user/subject",
		Handler:      education.GetUserSubjects,
		NeedToken:    true,
	},
}
