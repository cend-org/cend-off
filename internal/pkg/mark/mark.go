package mark

import (
	"context"
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/pkg/user/authorization"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
)

func RateUser(ctx context.Context, input *model.MarkInput) (*model.Mark, error) {
	var (
		tok         *token.Token
		studentMark model.Mark
		err         error
	)

	studentMark = model.MapMarkInputToMark(*input, studentMark)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &studentMark, errx.UnAuthorizedError
	}

	if authorization.IsUserStudent(tok.UserID) {
		return &studentMark, errx.UnAuthorizedError
	}

	if authorization.IsUserParent(tok.UserID) {
		return &studentMark, errx.UnAuthorizedError
	}

	if studentMark.AuthorMark > 5 {
		return &studentMark, errx.Lambda(errors.New("value exceed 5 star"))
	}

	studentMark.AuthorID = tok.UserID
	err = SetUserMark(studentMark)
	if err != nil {
		return &studentMark, errx.DbInsertError
	}

	return &studentMark, nil
}

func GetUserAverageMark(ctx context.Context, userID int) (*int, error) {
	var (
		userMarks   []model.Mark
		err         error
		totalMark   int
		averageMark int
	)

	err = database.GetMany(&userMarks, `SELECT user_mark.* FROM user_mark WHERE user_id = ?`, userID)
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

func GetUserMarkComment(ctx context.Context) ([]model.Mark, error) {
	var (
		tok  *token.Token
		err  error
		mark []model.Mark
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return mark, errx.UnAuthorizedError
	}

	if authorization.IsUserStudent(tok.UserID) || authorization.IsUserParent(tok.UserID) {
		return mark, errx.UnAuthorizedError
	}

	err = database.GetMany(&mark,
		`SELECT user_mark.* 
			FROM user_mark
			WHERE user_mark.author_id= ?;`, tok.UserID)
	if err != nil {
		return mark, errx.DbGetError
	}

	return mark, nil
}

/*
	UTILS
*/

func SetUserMark(userMark model.Mark) (err error) {
	_, err = database.InsertOne(userMark)
	if err != nil {
		return err
	}
	return nil
}
