package user

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/pkg/errors"
	"net/mail"
)

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
