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
	"github.com/pkg/errors"
	"net/mail"
	"strings"
)

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

	if usr.ID == 0 {
		usr, err = NewUserWithEmail(input)
		if err != nil {
			return nil, err
		}
	}

	authorization.UserID = usr.ID
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

func NewPassword(ctx context.Context, password string) (*bool, error) {
	var (
		psw  model.Password
		err  error
		done bool
		tok  *token.Token
	)

	if len(strings.TrimSpace(password)) == 0 {
		return nil, errors.New("password cannot be empty")
	}

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	psw.UserID = tok.UserId
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

func GetUserPasswordHash(userId uint) (hash string, err error) {
	return hash, err
}

func LogIn(ctx context.Context, email string, password string) (*string, error) {
	var access string
	var err error

	access, err = LoginWithEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}

	return &access, err
}

func GetUserAuthorizationLink(ctx context.Context, id int) (*model.UserAuthorizationLink, error) {
	panic(fmt.Errorf("not implemented: UserAuthorizationLink - userAuthorizationLink"))
}

func GetUserAuthorizationLinks(ctx context.Context) ([]model.UserAuthorizationLink, error) {
	panic(fmt.Errorf("not implemented: UserAuthorizationLink - userAuthorizationLink"))
}

func GetPasswordHistory(ctx context.Context) ([]model.Password, error) {
	panic(fmt.Errorf("not implemented: GetPasswordHistory - getPasswordHistory"))
}

func ActivateUser(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: ActivateUser - activateUser"))
}

func MyProfile(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Register - Register"))
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

	usr.ID = int(id)

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

	err = database.Get(&psw, `SELECT * FROM password WHERE user_id = ? ORDER BY created_at DESC  LIMIT 1`, usr.ID)
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
