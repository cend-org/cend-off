package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/authentication"
	"github.com/cend-org/duval/internal/database"
	pwd "github.com/cend-org/duval/internal/password"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/code"
	"github.com/cend-org/duval/pkg/user/authorization"
	"github.com/pkg/errors"
	"net/mail"
	"strings"
	"time"
)

const (
	StatusNew        = 0
	StatusUnverified = 1

	StatusNeedPassword = 2

	StatusOnboardingInProgress = 3

	StatusActive = 4
)

func Register(ctx context.Context, input *model.UserInput, as int) (*string, error) {
	var (
		user               model.User
		err                error
		tokenStr           string
		authorizationLevel int
	)
	user = model.MapUserInputToUser(*input, user)
	authorizationLevel = as

	if !utils.IsValidEmail(user.Email) {
		return &tokenStr, errx.InvalidEmailError
	}

	_, err = GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &tokenStr, errx.DbGetError
	}

	if user.Id > state.ZERO {
		return &tokenStr, errx.DuplicateUserError
	}

	user.Matricule, err = utils.GenerateMatricule()
	if err != nil {
		return &tokenStr, errx.GenerateMatriculeError
	}

	if user.Name == state.EMPTY {
		user.Name = user.Matricule
	}

	if user.NickName == state.EMPTY {
		user.NickName = user.Matricule
	}

	if input.BirthDate != nil {
		user.Age = ComputeAge(user.BirthDate)
	}

	userId, err := database.InsertOne(user)
	if err != nil {
		return &tokenStr, errx.DuplicateUserError
	}

	user.Id = userId

	err = authorization.NewUserAuthorization(user.Id, authorizationLevel)
	if err != nil {
		return &tokenStr, errx.AuthorizationError
	}

	tokenStr, err = LoginWithEmail(user.Email)
	if err != nil {
		return nil, err
	}

	err = code.NewUserVerificationCode(user.Id)
	if err != nil {
		return &tokenStr, errx.VerificationError
	}

	return &tokenStr, nil
}

func RegisterWithEmail(ctx context.Context, input string, as int) (*string, error) {
	var (
		usr           model.User
		authorization model.Authorization
		pat           string
		err           error
	)

	usr, err = GetUserByEmail(input)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if usr.Id == state.ZERO {
		usr, err = NewUserWithEmail(input)
		if err != nil {
			return nil, err
		}
	}

	authorization.UserId = usr.Id
	authorization.Level = as

	_, err = database.Insert(authorization)
	if err != nil {
		return nil, err
	}

	pat, err = LoginWithEmail(input)
	if err != nil {
		return nil, err
	}

	return &pat, err
}

func GetUserPasswordHash(userId uint) (hash string, err error) {
	return hash, err
}

func LogIn(ctx context.Context, email string, password string) (*string, error) {
	var access string
	var err error

	access, err = LoginWithEmailAndPassword(email, password)
	if err != nil {
		return nil, errx.AuthError
	}

	return &access, nil
}

func GetUserAuthorizationLink(ctx context.Context, id int) (*model.UserAuthorizationLink, error) {
	panic(fmt.Errorf("not implemented: UserAuthorizationLink - userAuthorizationLink"))
}

func GetUserAuthorizationLinks(ctx context.Context) ([]model.UserAuthorizationLink, error) {
	panic(fmt.Errorf("not implemented: UserAuthorizationLink - userAuthorizationLink"))
}

func NewPassword(ctx context.Context, password string) (*bool, error) {
	var (
		psw  model.Password
		err  error
		done bool
		tok  *token.Token
	)

	if len(strings.TrimSpace(password)) == state.ZERO {
		return nil, errors.New("password cannot be empty")
	}

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	psw.UserId = tok.UserId
	psw.Hash = pwd.Encode(password)
	if err != nil {
		return nil, err
	}

	_, err = database.Insert(psw)
	if err != nil {
		return nil, err
	}

	done = true

	return &done, err
}

func GetPasswordHistory(ctx context.Context) ([]model.Password, error) {
	var (
		passwords []model.Password
		tok       *token.Token
		err       error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return passwords, errx.UnAuthorizedError
	}

	err = database.Select(&passwords,
		`SELECT password.*
			FROM password
			WHERE password.user_id = ?
			ORDER BY password.created_at DESC`, tok.UserId)
	if err != nil {
		return passwords, errx.DbGetError
	}

	return passwords, nil
}

func ActivateUser(ctx context.Context) (*model.User, error) {
	var (
		tok *token.Token
		err error
		usr model.User
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &usr, errx.UnAuthorizedError
	}

	usr, err = GetUserWithId(tok.UserId)
	if err != nil {
		return &usr, errx.DbGetError
	}

	if usr.Status <= StatusNeedPassword {
		return &usr, errx.NeedPasswordError
	}

	usr.Status = StatusActive

	return &usr, nil
}

func MyProfile(ctx context.Context) (*model.User, error) {
	var (
		tok  *token.Token
		err  error
		user model.User
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &user, errx.UnAuthorizedError
	}

	user, err = GetUserWithId(tok.UserId)
	if err != nil {
		return &user, errx.DbGetError
	}

	return &user, nil
}

func UpdMyProfile(ctx context.Context, input *model.UserInput) (*string, error) {
	var (
		err  error
		tok  *token.Token
		usr  model.User
		done string
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &done, errx.UnAuthorizedError
	}

	usr, err = GetUserWithId(tok.UserId)
	if err != nil {
		return &done, errx.DbGetError
	}

	usr = model.MapUserInputToUser(*input, usr)

	//To do
	err = database.Update(usr)
	if err != nil {
		return &done, errx.DbUpdateError
	}

	done = "success"
	return &done, nil
}

/*
	UTILS
*/

func GetUserByEmail(email string) (usr model.User, err error) {
	err = database.Get(&usr, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return usr, err
	}
	return usr, err
}

func IsEmailValid(email string) (err error) {
	_, err = mail.ParseAddress(email)
	if err != nil {
		return err
	}
	return err
}

func NewUserWithEmail(email string) (usr model.User, err error) {
	isEmailValid := IsEmailValid(email) != nil
	if isEmailValid {
		return usr, errors.Wrap(err, "email is not valid")
	}

	usr.Email = email
	id, err := database.Insert(usr)
	if err != nil {
		return usr, err
	}

	usr.Id = int(id)

	return usr, err
}

/*
	UTILS
*/

func LoginWithEmail(email string) (access string, err error) {
	var usr model.User

	err = database.Get(&usr, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return
	}

	access, err = authentication.NewAccessToken(usr)
	if err != nil {
		return access, err
	}

	return access, err
}

func LoginWithEmailAndPassword(email, pass string) (access string, err error) {
	var usr model.User
	var psw model.Password

	usr, err = GetUserByEmail(email)
	if err != nil {
		return access, err
	}

	err = database.Get(&psw, `SELECT * FROM password WHERE user_id = ? ORDER BY created_at DESC  LIMIT 1`, usr.Id)
	if err != nil {
		return access, err
	}

	if !pwd.ValidatePassword(psw.Hash, pass) {
		return access, errors.New("email and password doesn't match any account")
	}

	access, err = authentication.NewAccessToken(usr)
	if err != nil {
		return access, err
	}

	return access, err
}

func GetUserWithId(id int) (user model.User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, id)
	if err != nil {
		return user, err
	}

	return user, err
}

func ComputeAge(birthDate time.Time) int {
	return time.Now().Year() - birthDate.Year()
}
