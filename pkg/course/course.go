package course

import (
	"context"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
)

func SetUserCoursePreference(ctx context.Context, isOnline bool) (*model.UserCoursePreference, error) {
	var (
		userCoursePreference model.UserCoursePreference
		err                  error
		tok                  *token.Token
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &userCoursePreference, errx.UnAuthorizedError
	}

	userCoursePreference.UserId = tok.UserId
	userCoursePreference.IsOnline = isOnline

	_, err = database.InsertOne(userCoursePreference)
	if err != nil {
		return &userCoursePreference, errx.DbInsertError
	}
	return &userCoursePreference, nil
}

func UpdUserCoursePreference(ctx context.Context, isOnline bool) (*model.UserCoursePreference, error) {
	var (
		userCoursePreference model.UserCoursePreference
		err                  error
		tok                  *token.Token
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &userCoursePreference, errx.UnAuthorizedError
	}

	userCoursePreference, err = GetUserCourse(tok.UserId)
	if err != nil {
		return &userCoursePreference, errx.DbInsertError
	}

	userCoursePreference.IsOnline = isOnline

	err = database.Update(userCoursePreference)
	if err != nil {
		return &userCoursePreference, errx.DbUpdateError
	}
	return &userCoursePreference, nil
}

func GetUserCoursePreference(ctx context.Context) (*model.UserCoursePreference, error) {
	var (
		userCoursePreference model.UserCoursePreference
		err                  error
		tok                  *token.Token
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &userCoursePreference, errx.UnAuthorizedError
	}

	userCoursePreference, err = GetUserCourse(tok.UserId)
	if err != nil {
		return &userCoursePreference, errx.DbGetError
	}
	return &userCoursePreference, nil
}

func GetUserCourse(userId int) (ucp model.UserCoursePreference, err error) {
	err = database.Get(&ucp, `SELECT * FROM user_course_preference WHERE  user_id = ? `, userId)
	if err != nil {
		return ucp, err
	}
	return ucp, nil
}
