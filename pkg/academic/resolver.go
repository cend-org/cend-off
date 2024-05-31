package academic

import (
	"context"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/user/authorization"
)

type AcademicQuery struct{}
type AcademicMutation struct{}

/*

	EDUCATIONS

*/

func (r *AcademicQuery) AcademicLevels(ctx context.Context) ([]model.AcademicLevel, error) {
	return GetAcademicLevels()
}

func (r *AcademicQuery) AcademicCourses(ctx context.Context, academicLevelID int) ([]model.AcademicCourse, error) {
	return GetAcademicCourses(academicLevelID)
}

func (r *AcademicMutation) NewUserAcademicCourses(ctx context.Context, courses []*model.UserAcademicCourseInput) (*bool, error) {
	var tok *token.Token
	var err error

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	return NewUserAcademicCourses(tok.UserId, courses)
}

func (r *AcademicMutation) NewStudentAcademicCoursesByParent(ctx context.Context, courses []*model.UserAcademicCourseInput, studentID int) (*bool, error) {
	var tok *token.Token
	var err error

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if !IsStudentParentLinked(tok.UserId, studentID) {
		return nil, errx.UlError
	}

	return NewUserAcademicCourses(studentID, courses)
}

/*

	LEVELS

*/

func (r *AcademicQuery) MultipleLevelAcademicCourses(ctx context.Context, academicLevelID []int) ([]model.AcademicCourse, error) {
	var courses []model.AcademicCourse

	for _, academicId := range academicLevelID {
		academicCourse, err := GetAcademicCourses(academicId)
		if err != nil {
			return nil, errx.SupportError
		}

		for _, course := range academicCourse {
			courses = append(courses, course)
		}
	}

	return courses, nil
}

func (r *AcademicQuery) UserAcademicLevels(ctx context.Context) ([]model.AcademicLevel, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	academicLevel, err := GetUserAcademicLevels(tok.UserId)
	if err != nil {
		return nil, err
	}
	return academicLevel, nil
}

func (r *AcademicQuery) StudentAcademicLevel(ctx context.Context, studentID int) ([]model.AcademicLevel, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if !IsStudentParentLinked(tok.UserId, studentID) {
		return nil, errx.UlError
	}

	academicLevel, err := GetUserAcademicLevels(studentID)
	if err != nil {
		return nil, err
	}
	return academicLevel, nil
}

func (r *AcademicMutation) SetUserAcademicLevel(ctx context.Context, academicLevelID int) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	err = SetUserAcademicLevel(tok.UserId, academicLevelID)
	if err != nil {
		return nil, err
	}

	status = true
	return &status, nil
}

func (r *AcademicMutation) SetStudentAcademicLevelByParent(ctx context.Context, academicLevelID int, studentID int) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if !IsStudentParentLinked(tok.UserId, studentID) {
		return &status, errx.UlError
	}
	err = SetUserAcademicLevel(studentID, academicLevelID)
	if err != nil {
		return nil, err
	}

	status = true
	return &status, nil
}

func (r *AcademicMutation) NewUserAcademicLevels(ctx context.Context, academicLevelIds []*int) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	err = SetUserAcademicLevels(tok.UserId, academicLevelIds)
	if err != nil {
		return nil, err
	}
	status = true
	return &status, nil
}

/*

	PREFERENCES

*/
// PREFERRED COURSE

func (r *AcademicQuery) UserCoursePreferences(ctx context.Context, userID int) ([]model.AcademicCourse, error) {
	var (
		tok     *token.Token
		err     error
		courses []model.AcademicCourse
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnknownLevelError
	}

	courses, err = GetUserPreferredCourse(userID)
	if err != nil {
		return nil, errx.SupportError
	}

	return courses, nil
}

func (r *AcademicQuery) CoursePreferences(ctx context.Context) ([]model.AcademicCourse, error) {
	var (
		tok     *token.Token
		err     error
		courses []model.AcademicCourse
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	courses, err = GetUserPreferredCourse(tok.UserId)
	if err != nil {
		return nil, errx.SupportError
	}

	return courses, nil
}

//IsOnline

func (r *AcademicQuery) UserPreferences(ctx context.Context, studentID int) (*model.UserAcademicCoursePreference, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if authorization.IsUserStudent(studentID) && !IsStudentParentLinked(tok.UserId, studentID) {
		return nil, errx.UlError
	}

	course, err := GetPreferences(studentID)
	if err != nil {
		return nil, errx.MissingPreferenceError
	}
	return &course, nil
}

func (r *AcademicQuery) Preferences(ctx context.Context) (*model.UserAcademicCoursePreference, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	course, err := GetPreferences(tok.UserId)
	if err != nil {
		return nil, errx.MissingPreferenceError
	}
	return &course, nil
}

func (r *AcademicMutation) UpdAcademicCoursePreference(ctx context.Context, coursesPreferences model.UserAcademicCoursePreferenceInput) (*model.UserAcademicCoursePreference, error) {
	var tok *token.Token
	var err error

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	preference, err := UpdStudentAcademicCoursesPreferenceByParent(tok.UserId, coursesPreferences)
	if err != nil {
		return nil, errx.SupportError
	}

	return &preference, nil
}

func (r *AcademicMutation) UpdStudentAcademicCoursesPreferenceByParent(ctx context.Context, coursesPreferences model.UserAcademicCoursePreferenceInput, studentID int) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if !IsStudentParentLinked(tok.UserId, studentID) {
		return nil, errx.UlError
	}

	_, err = UpdStudentAcademicCoursesPreferenceByParent(studentID, coursesPreferences)
	if err != nil {
		return nil, errx.SupportError
	}
	status = true

	return &status, nil
}

/*

	LINKS

*/
//TUTOR

func (r *AcademicQuery) SuggestTutor(ctx context.Context, studentID int) (*model.User, error) {
	var (
		tok  *token.Token
		err  error
		user model.User
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if !IsStudentParentLinked(tok.UserId, studentID) {
		return nil, errx.UlError
	}

	user, err = GetTutorWithPreferredCourse(studentID)
	if err != nil {
		return nil, err
	}

	if user.Id == state.ZERO {
		return nil, errx.EmptyTutorError
	}

	return &user, nil
}

func (r *AcademicQuery) SuggestOtherTutor(ctx context.Context, studentID int, lastTutorID int) (*model.User, error) {
	var (
		tok  *token.Token
		err  error
		user model.User
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if !IsStudentParentLinked(tok.UserId, studentID) {
		return nil, errx.UlError
	}

	user, err = GetOtherTutorWithPreferredCourse(studentID, lastTutorID)
	if err != nil {
		return nil, err
	}

	if user.Id == state.ZERO {
		return nil, errx.EmptyTutorError
	}

	return &user, nil
}

func (r *AcademicQuery) SuggestTutorToUser(ctx context.Context) (*model.User, error) {
	var (
		tok  *token.Token
		err  error
		user model.User
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	user, err = GetTutorWithPreferredCourse(tok.UserId)
	if err != nil {
		return nil, err
	}

	if user.Id == state.ZERO {
		return nil, errx.EmptyTutorError
	}

	return &user, nil
}

func (r *AcademicQuery) SuggestOtherTutorToUser(ctx context.Context, lastTutorID int) (*model.User, error) {
	var (
		tok  *token.Token
		err  error
		user model.User
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	user, err = GetOtherTutorWithPreferredCourse(tok.UserId, lastTutorID)
	if err != nil {
		return nil, err
	}

	if user.Id == state.ZERO {
		return nil, errx.EmptyTutorError
	}

	return &user, nil
}

func (r *AcademicMutation) NewStudentTutorByParent(ctx context.Context, tutorID int, studentID int) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	if !IsStudentParentLinked(tok.UserId, studentID) {
		return nil, errx.UlError
	}

	_, err = AddStudentToTutor(tutorID, studentID)
	if err != nil {
		return nil, err
	}

	status = true
	return &status, nil
}

func (r *AcademicMutation) NewStudentTutor(ctx context.Context, userID int) (*model.User, error) {
	var (
		tok     *token.Token
		err     error
		student *model.User
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if !authorization.IsUserStudent(tok.UserId) && !authorization.IsUserTutor(tok.UserId) {
		return nil, errx.UnAuthorizedError
	}

	if !authorization.IsUserStudent(userID) && !authorization.IsUserTutor(userID) {
		return nil, errx.UnAuthorizedError
	}

	if authorization.IsUserTutor(tok.UserId) && authorization.IsUserStudent(userID) {
		student, err = AddStudentToTutor(tok.UserId, userID)
		if err != nil {
			return nil, err
		}
	}

	if authorization.IsUserTutor(userID) && authorization.IsUserStudent(tok.UserId) {
		student, err = AddStudentToTutor(userID, tok.UserId)
		if err != nil {
			return nil, err
		}
	}

	return student, nil
}

// PROFESSOR

func (r *AcademicQuery) ProfessorStudent(ctx context.Context, keyWord string) ([]model.User, error) {
	var (
		tok  *token.Token
		err  error
		user []model.User
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if !authorization.IsUserProfessor(tok.UserId) {
		return nil, errx.UnAuthorizedError
	}

	user, err = GetProfessorStudent(keyWord)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AcademicMutation) NewStudentProfessor(ctx context.Context, userID int) (*model.User, error) {
	var (
		tok     *token.Token
		err     error
		student *model.User
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if !authorization.IsUserStudent(tok.UserId) && !authorization.IsUserProfessor(tok.UserId) {
		return nil, errx.UnAuthorizedError
	}

	if !authorization.IsUserStudent(userID) && !authorization.IsUserProfessor(userID) {
		return nil, errx.UnAuthorizedError
	}

	if authorization.IsUserProfessor(tok.UserId) && authorization.IsUserStudent(userID) {
		student, err = AddStudentToProfessor(tok.UserId, userID)
		if err != nil {
			return nil, err
		}
	}

	if authorization.IsUserProfessor(userID) && authorization.IsUserStudent(tok.UserId) {
		student, err = AddStudentToProfessor(userID, tok.UserId)
		if err != nil {
			return nil, err
		}
	}

	return student, nil
}

// STUDENT

func (r *AcademicMutation) UserStudent(ctx context.Context, name string, familyName string) (*model.User, error) {
	var (
		tok     *token.Token
		err     error
		student *model.User
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	student, err = AddStudentToLink(tok.UserId, name, familyName)
	if err != nil {
		return nil, err
	}

	return student, nil
}
