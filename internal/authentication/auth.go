package authentication

import (
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/password"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/user"
	"time"
)

func NewAccessToken(usr model.User) (string, error) {
	var (
		auth   model.Authorization
		tok    token.Token
		access string
		err    error
	)

	err = database.Get(&auth, `SELECT * FROM authorization WHERE user_id = ? ORDER BY access_level desc limit 1`, usr.ID)
	if err != nil {
		return "", err
	}

	/* fill token field */
	tok.UserId = usr.ID
	tok.AccessLevel = auth.AccessLevel

	tok.ExpirationDate.Value = time.Now().Add(time.Hour * 24)

	/* parse token */
	access, err = token.New(tok)
	if err != nil {
		return access, err
	}

	time.Sleep(time.Second * 2)

	return access, err
}

func LoginWithEmail(email string) (access string, err error) {
	var usr model.User

	err = database.Get(&usr, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return
	}

	access, err = NewAccessToken(usr)
	if err != nil {
		return access, err
	}

	return access, err
}

func LoginWithEmailAndPassword(email, pass string) (access string, err error) {
	var usr model.User
	var psw model.Password

	usr, err = user.GetUserByEmail(email)
	if err != nil {
		return access, err
	}

	err = database.Get(&psw, `SELECT * FROM password WHERE user_id = ? ORDER BY created_at DESC  LIMIT 1`, usr.ID)
	if err != nil {
		return access, err
	}

	if !password.ValidatePassword(psw.Hash, pass) {
		return access, errors.New("email and password doesn't match any account")
	}

	access, err = NewAccessToken(usr)
	if err != nil {
		return access, err
	}

	return access, err
}
