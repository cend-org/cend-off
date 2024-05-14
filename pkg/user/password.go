package user

import (
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/xorcare/pointer"
	"strings"
)

func NewPassword(userId int, new model.PasswordInput) (ret *bool, err error) {
	password := model.MapPasswordInputToPassword(new, model.Password{})

	if strings.TrimSpace(password.Hash) == state.EMPTY {
		return pointer.Bool(false), errors.New("password cannot be empty")
	}

	if userId == state.ZERO {
		return pointer.Bool(false), errors.New("user must be set")
	}

	if !utils.PasswordHasValidLength(password.Hash) {
		return pointer.Bool(false), errors.New("password length is not valid")
	}

	password.Hash, err = utils.HashPassword(password.Hash)
	if err != nil {
		return nil, errors.New("cannot protect your password. Something went wrong")
	}

	password.UserId = userId

	_, err = database.Insert(password)
	if err != nil {
		return nil, errors.New("cannot store your password")
	}

	return pointer.Bool(true), err
}

func getUserActivePassword(userId int) (password model.Password, err error) {
	err = database.Get(&password, `SELECT * FROM password WHERE user_id = ? ORDER BY created_at DESC LIMIT 1`, userId)
	if err != nil {
		return password, err
	}

	return password, err
}
