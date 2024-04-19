package user

import (
	"context"
	"database/sql"
	"duval/internal/authentication"
	"duval/internal/graph/model"
	"duval/internal/pkg/code"
	"duval/internal/pkg/user/authorization"
	"duval/internal/pkg/user/password"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"errors"
	"time"
)

const (
	StatusNew        = 0
	StatusUnverified = 1

	StatusNeedPassword = 2

	StatusOnboardingInProgress = 3

	StatusActive = 4
)

func Register(input *model.NewUserInput, userType *int) (string, error) {
	var (
		user               model.User
		err                error
		tokenStr           string
		authorizationLevel int
	)

	user = model.User{
		Name:       input.Name,
		FamilyName: input.FamilyName,
		NickName:   input.NickName,
		BirthDate:  input.BirthDate,
		Email:      input.Email,
		Sex:        input.Sex,
		Lang:       input.Lang,
	}
	authorizationLevel = *userType

	if !utils.IsValidEmail(user.Email) {
		return tokenStr, errx.InvalidEmailError
	}

	_, err = GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return tokenStr, errx.Lambda(err)
	}

	if user.Id > state.ZERO {
		return tokenStr, errx.DuplicateUserError
	}

	user.Matricule, err = utils.GenerateMatricule()
	if err != nil {
		return tokenStr, errx.Lambda(err)
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

func MyProfile(ctx *context.Context) (*model.User, error) {
	var (
		tok  *authentication.Token
		err  error
		user model.User
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &user, errx.Lambda(err)
	}

	user, err = GetUserWithId(tok.UserId)
	if err != nil {
		return &user, errx.Lambda(err)
	}

	return &user, nil
}

func GetCode(ctx *context.Context) (*model.Code, error) {
	var (
		err            error
		tok            *authentication.Token
		validationCode model.Code
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &validationCode, errx.Lambda(err)
	}

	err = database.Get(&validationCode, `SELECT * FROM code WHERE user_id = ? ORDER BY created_at desc LIMIT 1`, tok.UserId)
	if err != nil {
		return &validationCode, errx.UnAuthorizedError
	}

	return &validationCode, nil
}

func VerifyUserEmailValidationCode(ctx *context.Context, validationCode int) (int, error) {
	var (
		err  error
		tok  *authentication.Token
		user model.User
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return 0, errx.UnAuthorizedError
	}

	err = code.IsUserVerificationCodeValid(tok.UserId, validationCode)
	if err != nil {
		return 0, errx.Lambda(err)
	}

	user, err = GetUserWithId(tok.UserId)
	if err != nil {
		return 0, errx.Lambda(err)
	}

	if user.Status < StatusNeedPassword {
		user.Status = StatusNeedPassword
	}

	err = database.Update(user)
	if err != nil {
		return 0, errx.Lambda(err)
	}

	return validationCode, nil
}

func SendUserEmailValidationCode(ctx *context.Context) (*model.User, error) {
	var (
		err  error
		tok  *authentication.Token
		user model.User
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &user, errx.UnAuthorizedError
	}

	err = code.NewUserVerificationCode(tok.UserId)
	if err != nil {
		return &user, errx.Lambda(err)

	}

	user, err = GetUserWithId(tok.UserId)
	if err != nil {
		return &user, errx.Lambda(err)
	}

	user.Status = StatusUnverified

	err = database.Update(user)
	if err != nil {
		return &user, errx.Lambda(err)
	}

	return &user, nil
}

func UpdMyProfile(ctx *context.Context, input map[string]interface{}) (*model.User, error) {
	var (
		err error
		tok *authentication.Token
		usr model.User
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &usr, errx.UnAuthorizedError
	}

	usr.Id = tok.UserId

	usr, err = GetUserWithId(usr.Id)
	if err != nil {
		return &usr, errx.DbGetError
	}
	time.Sleep(100)
	err = utils.ApplyChanges(input, &usr)
	if err != nil {
		return &usr, errx.Lambda(err)
	}
	err = database.Update(usr)
	if err != nil {
		return &usr, errx.Lambda(err)
	}

	return &usr, nil
}

func Login(input *model.UserLogin) (string, error) {
	var (
		err      error
		usr      model.User
		tokenStr string
	)

	// GET USER DATA
	usr, err = GetUserByEmail(input.Email)
	if err != nil {
		return tokenStr, errx.Lambda(err)
	}

	if !password.IsPasswordValid(usr.Id, input.Password) {
		return tokenStr, errx.Lambda(err)
	}

	tokenStr, err = authentication.GetTokenString(usr.Id)
	if err != nil {
		return tokenStr, errx.Lambda(err)

	}

	return tokenStr, nil
}

func NewPassword(ctx *context.Context, input *model.NewPassword) (*string, error) {
	var (
		err    error
		tok    *authentication.Token
		usr    model.User
		status string
		pass   model.Password
	)

	pass.Psw = input.Psw
	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	err = password.CreatePassword(tok.UserId, pass)
	if err != nil {
		return &status, errx.Lambda(err)
	}

	usr, err = GetUserWithId(tok.UserId)
	if err != nil {
		return &status, errx.Lambda(err)
	}

	if usr.Status < StatusOnboardingInProgress {
		usr.Status = StatusOnboardingInProgress
		err = database.Update(usr)
		if err != nil {
			return &status, errx.Lambda(err)
		}
	}
	status = "success"
	return &status, nil
}

func GetPasswordHistory(ctx *context.Context) ([]*model.Password, error) {
	var (
		passwords    []model.Password
		gqlPasswords []*model.Password
		tok          *authentication.Token
		err          error
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return gqlPasswords, errx.UnAuthorizedError
	}
	err = database.GetMany(&passwords,
		`SELECT password.*
			FROM password
			WHERE password.user_id = ?
			ORDER BY password.created_at DESC`, tok.UserId)
	if err != nil {
		return gqlPasswords, errx.DbGetError
	}

	for _, pass := range passwords {
		gqlPasswords = append(gqlPasswords, &model.Password{
			Id:          pass.Id,
			CreatedAt:   pass.CreatedAt,
			UpdatedAt:   pass.UpdatedAt,
			DeletedAt:   pass.DeletedAt,
			UserId:      pass.UserId,
			Psw:         pass.Psw,
			ContentHash: pass.ContentHash,
		})
	}

	return gqlPasswords, nil
}

func ActivateUser(ctx *context.Context) (*model.User, error) {
	var (
		tok *authentication.Token
		err error
		usr model.User
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &usr, errx.Lambda(err)
	}

	usr, err = GetUserWithId(tok.UserId)
	if err != nil {
		return &usr, errx.Lambda(err)
	}

	if usr.Status <= StatusNeedPassword {
		return &usr, errx.Lambda(err)
	}

	usr.Status = StatusActive

	return &usr, nil
}

func RegisterByEmail(authorizationLevel *int, email *string) (*string, error) {
	var (
		user     model.User
		err      error
		tokenStr string
	)

	user.Email = *email
	if !utils.IsValidEmail(user.Email) {
		return &tokenStr, errx.InvalidEmailError
	}

	_, err = GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &tokenStr, errx.Lambda(err)
	}

	if user.Id > state.ZERO {
		return &tokenStr, errx.DuplicateUserError
	}

	user.Matricule, err = utils.GenerateMatricule()
	if err != nil {
		return &tokenStr, errx.Lambda(err)
	}

	if user.Name == state.EMPTY {
		user.Name = user.Matricule
	}

	if user.NickName == state.EMPTY {
		user.NickName = user.Matricule
	}

	user.Id, err = database.InsertOne(user)
	if err != nil {
		return &tokenStr, errx.DbInsertError
	}

	err = authorization.NewUserAuthorization(user.Id, uint(*authorizationLevel))
	if err != nil {
		return &tokenStr, errx.Lambda(err)
	}

	tokenStr, err = authentication.GetTokenString(user.Id)
	if err != nil {
		return &tokenStr, errx.Lambda(err)
	}

	err = code.NewUserVerificationCode(user.Id)
	if err != nil {
		return &tokenStr, errx.Lambda(err)
	}

	return &tokenStr, nil
}

/*

	UTILITIES

*/

func GetUserByEmail(email string) (user model.User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return user, err
	}

	return user, err
}

func GetUserWithId(id uint) (user model.User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, id)
	if err != nil {
		return user, err
	}

	return user, err
}
