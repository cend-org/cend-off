package education

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

func GetSubjects(eduId int) ([]*model.Subject, error) {
	var (
		err      error
		subjects []*model.Subject
	)

	err = database.Select(&subjects, `SELECT * FROM subject WHERE education_level_id = ?`, eduId)
	if err != nil {
		return subjects, errx.Lambda(err)
	}

	return subjects, nil
}

func GetEducation() ([]*model.Education, error) {
	var (
		err  error
		edus []*model.Education
	)

	err = database.Select(&edus, `SELECT * FROM education WHERE id > 0 ORDER BY  created_at`)
	if err != nil {
		return edus, errx.Lambda(err)
	}

	return edus, nil
}

// User educationLevel

func SetUserEducationLevel(ctx *context.Context, subject *model.SubjectInput) (*model.Education, error) {
	var (
		tok                       *authentication.Token
		userEducationLevelSubject model.UserEducationLevelSubject
		err                       error
		userLevel                 model.Education
	)
	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &userLevel, errx.UnAuthorizedError
	}
	if tok.UserId == state.ZERO {
		return &userLevel, errx.UnAuthorizedError
	}

	if !authorization.IsUserStudent(tok.UserId) {
		if !authorization.IsUserProfessor(tok.UserId) {
			return &userLevel, errx.Lambda(errors.New("not a user or a professor"))
		}
	}

	//err = ctx.ShouldBindJSON(&subject)
	//if err != nil {
	//	return &userLevel, errx.ParseError
	//}

	userEducationLevelSubject.SubjectId = uint(subject.Id)
	userEducationLevelSubject.UserId = tok.UserId

	_, err = database.InsertOne(userEducationLevelSubject)
	if err != nil {
		return &userLevel, errx.DbInsertError
	}

	userLevel, err = GetUserLevel(tok.UserId)
	if err != nil {
		return &userLevel, errx.DbGetError
	}

	return &userLevel, nil
}

func GetUserEducationLevel(ctx *context.Context) (*model.Education, error) {
	var (
		err       error
		tok       *authentication.Token
		userLevel model.Education
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &userLevel, errx.UnAuthorizedError
	}

	userLevel, err = GetUserLevel(tok.UserId)
	if err != nil {
		return &userLevel, errx.DbGetError

	}
	return &userLevel, nil

}

func UpdateUserEducationLevel(ctx *context.Context, subject *model.SubjectInput) (*model.Education, error) {
	var (
		tok                              *authentication.Token
		currentUserEducationLevelSubject model.UserEducationLevelSubject
		userEducationLevelSubject        model.UserEducationLevelSubject
		err                              error
		userLevel                        model.Education
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &userLevel, errx.UnAuthorizedError
	}
	if tok.UserId == state.ZERO {
		return &userLevel, errx.UnAuthorizedError
	}

	if !authorization.IsUserStudent(tok.UserId) {
		return &userLevel, errx.Lambda(errors.New("user is not a student"))
	}

	//err = ctx.ShouldBindJSON(&subject)
	//if err != nil {
	//	return &userLevel, errx.ParseError
	//}

	err = database.Get(&currentUserEducationLevelSubject, `SELECT user_education_level_subject.* FROM user_education_level_subject
			WHERE user_education_level_subject.user_id = ?`, tok.UserId)
	if err != nil {
		return &userLevel, errx.DbGetError
	}

	err = RemoveUserEducationLevelSubject(currentUserEducationLevelSubject)
	if err != nil {
		return &userLevel, errx.DbDeleteError
	}

	userEducationLevelSubject.SubjectId = subject.Id
	userEducationLevelSubject.UserId = tok.UserId
	_, err = database.InsertOne(userEducationLevelSubject)
	if err != nil {
		return &userLevel, errx.DbInsertError
	}

	userLevel, err = GetUserLevel(tok.UserId)
	if err != nil {
		return &userLevel, errx.DbGetError
	}

	return &userLevel, nil
}

// User education subjects

func GetUserSubjects(ctx *context.Context) ([]*model.Subject, error) {
	var (
		subjects []*model.Subject
		tok      *authentication.Token
		err      error
	)
	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return subjects, errx.UnAuthorizedError
	}

	if authorization.IsUserParent(tok.UserId) || authorization.IsUserTutor(tok.UserId) {
		return subjects, errx.UnAuthorizedError
	}

	if authorization.IsUserStudent(tok.UserId) {
		err = database.GetMany(&subjects, `SELECT subject.* 
											FROM subject
											WHERE subject.education_level_id = (SELECT education.id FROM education  JOIN subject ON education.id  =  subject.education_level_id JOIN user_education_level_subject ON subject.id = user_education_level_subject.subject_id
                                   			WHERE user_education_level_subject.user_id = ?)`, tok.UserId)
		if err != nil {
			return subjects, errx.DbGetError
		}
	}

	if authorization.IsUserProfessor(tok.UserId) {
		err = database.GetMany(&subjects, `SELECT subject.* FROM subject
			JOIN user_education_level_subject  ON subject.id = user_education_level_subject.subject_id
			WHERE user_education_level_subject.user_id = ?`, tok.UserId)
		if err != nil {
			return subjects, errx.DbGetError
		}
	}

	return subjects, nil
}

/*
	UTILS
*/

func GetUserLevel(userId uint) (educationLevel model.Education, err error) {

	err = database.Get(&educationLevel,
		`SELECT education.* FROM education
				JOIN subject ON education.id  =  subject.education_level_id
				JOIN user_education_level_subject ON subject.id = user_education_level_subject.subject_id
			WHERE user_education_level_subject.user_id = ?`, userId)
	if err != nil {
		return educationLevel, err
	}

	return educationLevel, err
}

func RemoveUserEducationLevelSubject(userEducationLevelSubject model.UserEducationLevelSubject) (err error) {
	err = database.Delete(userEducationLevelSubject)
	if err != nil {
		return err
	}
	return err
}
