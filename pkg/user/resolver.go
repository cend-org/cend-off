package user

import (
	"context"
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/media/cover"
	"github.com/cend-org/duval/pkg/media/cv"
	"github.com/cend-org/duval/pkg/media/profile"
	"github.com/cend-org/duval/pkg/media/video"
	"github.com/xorcare/pointer"
)

type UserQuery struct{}
type UserMutation struct{}

/*

	PROFILE

*/

func (r *UserMutation) UpdateUserStatus(ctx context.Context, rt int) (*bool, error) {
	var (
		tok    *token.Token
		user   model.User
		err    error
		status bool
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	user, err = GetUserWithId(tok.UserId)
	if err != nil {
		return &status, errx.SupportError
	}

	user.Status = rt

	err = database.Update(user)
	if err != nil {
		return &status, errx.SupportError
	}
	status = true
	return &status, nil
}

func (r *UserQuery) MyProfile(ctx context.Context) (*model.User, error) {
	var tok *token.Token
	var err error

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errors.New("unAuthorized")
	}

	return MyProfile(tok.UserId)
}

func (r *UserQuery) UserProfile(ctx context.Context, userID int) (*model.User, error) {
	var tok *token.Token
	var err error

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	if tok.UserId == state.ZERO {
		return nil, errx.UnknownUserError
	}

	return MyProfile(userID)
}

// NEW PROFILE

func (r *UserMutation) NewStudent(ctx context.Context, email string) (*model.BearerToken, error) {
	return NewStudent(email)
}

func (r *UserMutation) NewParent(ctx context.Context, email string) (*model.BearerToken, error) {
	return NewParent(email)
}

func (r *UserMutation) NewTutor(ctx context.Context, email string) (*model.BearerToken, error) {
	return NewTutor(email)
}

func (r *UserMutation) NewProfessor(ctx context.Context, email string) (*model.BearerToken, error) {
	return NewProfessor(email)
}

// UPDATE EXISTING PROFILE

func (r *UserMutation) UpdateMyProfile(ctx context.Context, profile model.UserInput) (*model.User, error) {
	var tok *token.Token
	var err error

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errors.New("unAuthorized")
	}

	user, err := UpdMyProfile(tok.UserId, profile)

	return user, nil
}

func (r *UserMutation) UpdateProfileAndPassword(ctx context.Context, profile model.UserInput, password model.PasswordInput) (*model.User, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errors.New("unAuthorized")
	}

	return UpdateProfileAndPassword(tok.UserId, profile, password)
}

func (r *UserMutation) UpdateStudentProfileByParent(ctx context.Context, profile model.UserInput, studentID int) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	if !IsStudentParentLinked(tok.UserId, studentID) {
		return &status, errx.UlError
	}

	err = UpdateStudent(studentID, profile)
	if err != nil {
		return &status, err
	}

	status = true
	return &status, nil
}

// AUTHENTICATION

func (r *UserMutation) NewPassword(ctx context.Context, password model.PasswordInput) (*bool, error) {
	var tok *token.Token
	var err error
	var status *bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return status, errors.New("unAuthorized")
	}

	status, err = NewPassword(tok.UserId, password)
	if err != nil {
		return status, err
	}

	return status, err
}

func (r *UserMutation) NewStudentsPassword(ctx context.Context, studentID int) (*string, error) {
	var tok *token.Token
	var err error
	var password string

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &password, errx.UnAuthorizedError
	}

	if !IsStudentParentLinked(tok.UserId, studentID) {
		return &password, errx.UlError
	}

	password, err = CreateStudentPassword(studentID)
	if err != nil {
		return &password, err
	}

	return &password, nil
}

func (r *UserMutation) Login(ctx context.Context, email string, password string) (*model.BearerToken, error) {
	return Login(email, password)
}

/*

	MEDIA

*/
// QUERY FOR ALL MEDIA

func (r *UserQuery) CoverLetter(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := cover.GetProfileLetter(tok.UserId)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}
	return &networkLink, nil
}

func (r *UserQuery) Cv(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := cv.GetProfileCv(tok.UserId)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}
	return &networkLink, nil
}

func (r *UserQuery) ProfileImage(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := profile.GetProfileImage(tok.UserId)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}
	return &networkLink, nil
}

func (r *UserQuery) VideoPresentation(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := video.GetProfileVideo(tok.UserId)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}
	return &networkLink, nil
}

// QUERY FOR ALL MEDIA THUMB

func (r *UserQuery) CoverLetterThumb(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := cover.GetProfileLetterThumb(tok.UserId)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}
	return &networkLink, nil
}

func (r *UserQuery) CvThumb(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := cv.GetProfileCvThumb(tok.UserId)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}
	return &networkLink, nil
}

func (r *UserQuery) ProfileImageThumb(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := profile.GetProfileImageThumb(tok.UserId)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}
	return &networkLink, nil
}

// QUERY FOR OTHER'S USER MEDIA

func (r *UserQuery) UserCoverLetter(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := cover.GetProfileLetter(userID)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}

	return &networkLink, nil
}

func (r *UserQuery) UserCv(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := cv.GetProfileCv(userID)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}

	return &networkLink, nil
}

func (r *UserQuery) UserProfileImage(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := profile.GetProfileImage(userID)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}

	return &networkLink, nil
}

func (r *UserQuery) UserVideoPresentation(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}

	networkLink, err := video.GetProfileVideo(userID)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}

	return &networkLink, nil
}

// QUERY FOR OTHER'S USER MEDIA THUMB

func (r *UserQuery) UserCoverLetterThumb(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := cover.GetProfileLetterThumb(userID)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}

	return &networkLink, nil
}

func (r *UserQuery) UserCvThumb(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := cv.GetProfileCvThumb(userID)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}

	return &networkLink, nil
}

func (r *UserQuery) UserProfileImageThumb(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := profile.GetProfileImageThumb(userID)
	if err != nil && errx.IsEmpty(err) {
		return nil, nil
	} else if err != nil && !errx.IsEmpty(err) {
		return &networkLink, errx.SupportError
	}

	return &networkLink, nil
}

// MUTATION FOR CLEANING MEDIA

func (r *UserMutation) RemoveCoverLetter(ctx context.Context) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	status, err = cover.RemoveProfileLetter(tok.UserId)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (r *UserMutation) RemoveCv(ctx context.Context) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	status, err = cv.RemoveProfileCv(tok.UserId)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (r *UserMutation) RemoveProfileImage(ctx context.Context) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	status, err = profile.RemoveProfileImage(tok.UserId)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (r *UserMutation) RemoveVideoPresentation(ctx context.Context) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	status, err = video.RemoveProfileVideo(tok.UserId)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (r *UserMutation) UpdateMyPassword(ctx context.Context, hash model.PasswordInput) (*bool, error) {
	var (
		tok    *token.Token
		err    error
		status bool
	)
	if hash.Hash == nil || !utils.PasswordHasValidLength(*hash.Hash) {
		return pointer.Bool(false), errx.PasswordLengthError
	}

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	status, err = UpdatePassword(tok.UserId, hash)
	if err != nil {
		return &status, err
	}

	return &status, nil
}

func (r *UserMutation) UpdateMyEmail(ctx context.Context, email string) (*model.User, error) {
	var (
		tok  *token.Token
		user model.User
		err  error
	)

	if email == state.EMPTY || !utils.IsValidEmail(email) {
		return &user, errx.InvalidEmailError
	}

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &user, errx.UnAuthorizedError
	}

	user, err = GetUserWithId(tok.UserId)
	if err != nil {
		return &user, errx.SupportError
	}

	user.Email = email

	return &user, nil
}
