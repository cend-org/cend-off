package mark

import (
	"context"
	"duval/internal/authentication"
	"duval/internal/graph/model"
	"duval/internal/pkg/user/authorization"
	"duval/internal/utils/errx"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"errors"
)

func RateUser(ctx *context.Context, mark *model.UserMarkInput) (*model.UserMark, error) {
	var (
		tok         *authentication.Token
		studentMark model.UserMark
		err         error
	)
	studentMark.UserId = mark.UserId
	studentMark.AuthorId = mark.AuthorId
	studentMark.AuthorMark = mark.AuthorMark
	studentMark.AuthorComment = mark.AuthorComment

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &studentMark, errx.UnAuthorizedError
	}

	if authorization.IsUserStudent(tok.UserId) {
		return &studentMark, errx.UnAuthorizedError
	}

	if authorization.IsUserParent(tok.UserId) {
		return &studentMark, errx.UnAuthorizedError
	}

	if studentMark.AuthorMark > 5 {
		return &studentMark, errx.Lambda(errors.New("value exceed 5 star"))
	}

	studentMark.AuthorId = tok.UserId
	err = SetUserMark(studentMark)
	if err != nil {
		return &studentMark, errx.DbInsertError
	}

	return &studentMark, nil
}

func GetUserAverageMark(ctx *context.Context, userId int) (*int, error) {
	var (
		userMarks   []model.UserMark
		err         error
		totalMark   uint
		averageMark int
	)

	err = database.GetMany(&userMarks, `SELECT user_mark.* FROM user_mark WHERE user_id = ?`, userId)
	if err != nil {
		return &averageMark, errx.DbGetError
	}

	totalAuthor := len(userMarks)
	if totalAuthor == state.ZERO {
		return &averageMark, errx.Lambda(errors.New("not rated "))
	}

	for _, userMark := range userMarks {
		totalMark = totalMark + userMark.AuthorMark
	}
	averageMark = int(totalMark) / (totalAuthor)

	return &averageMark, nil
}

func GetUserMarkComment(ctx *context.Context) ([]*model.UserMark, error) {
	var (
		tok  *authentication.Token
		err  error
		mark []*model.UserMark
	)
	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return mark, errx.UnAuthorizedError
	}

	if authorization.IsUserStudent(tok.UserId) || authorization.IsUserParent(tok.UserId) {
		return mark, errx.UnAuthorizedError
	}

	err = database.GetMany(&mark,
		`SELECT user_mark.* 
			FROM user_mark
			WHERE user_mark.author_id= ?;`, tok.UserId)
	if err != nil {
		return mark, errx.DbGetError
	}

	return mark, nil
}

//
/*
	UTILS
*/

func SetUserMark(userMark model.UserMark) (err error) {
	_, err = database.InsertOne(userMark)
	if err != nil {
		return err
	}
	return nil
}
