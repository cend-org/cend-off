package user

import (
	"database/sql"
	"duval/internal/authentication"
	"duval/internal/code"
	"duval/internal/graph/model"
	"duval/internal/pkg/user/authorization"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"errors"
)

const (
	StatusNew        = 0
	StatusUnverified = 1

	StatusNeedPassword = 2

	StatusOnboardingInProgress = 3

	StatusActive = 4
)

func CreateUser(input *model.NewUserInput, userType *int) (string, error) {
	var (
		user               model.User
		err                error
		tokenStr           string
		authorizationLevel int
	)

	user.Email = input.Email
	user.Name = input.Name
	user.FamilyName = input.FamilyName
	user.NickName = input.NickName
	user.BirthDate = input.BirthDate
	user.Sex = input.Sex
	user.Lang = input.Lang
	authorizationLevel = *userType

	if !utils.IsValidEmail(user.Email) {
		return "", errx.InvalidEmailError
	}

	_, err = GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", errx.Lambda(err)
	}

	if user.Id > state.ZERO {
		return "", errx.DuplicateUserError
	}

	user.Matricule, err = utils.GenerateMatricule()
	if err != nil {
		return "", errx.Lambda(err)
	}

	if user.Name == state.EMPTY {
		user.Name = user.Matricule
	}

	if user.NickName == state.EMPTY {
		user.NickName = user.Matricule
	}

	userId, err := database.InsertOne(user)
	if err != nil {
		return tokenStr, errx.DuplicateUserError
	}

	user.Id = userId

	err = authorization.NewUserAuthorization(user.Id, uint(authorizationLevel))
	if err != nil {
		return tokenStr, errx.Lambda(err)
	}

	tokenStr, err = authentication.GetTokenString(user.Id)
	if err != nil {
		return tokenStr, errx.Lambda(err)
	}

	err = code.NewUserVerificationCode(user.Id)
	if err != nil {
		return tokenStr, errx.Lambda(err)
	}

	return tokenStr, nil
}

/*
UTILS
*/

func GetUserByEmail(email string) (user model.User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return user, err
	}

	return user, err
}
