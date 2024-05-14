package user

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/authentication"
	"github.com/cend-org/duval/internal/utils"
	"github.com/pkg/errors"
)

func Login(email string, password string) (bearer *model.BearerToken, err error) {
	var user model.User
	var psw model.Password
	var T string

	user, err = getUserByEmail(email)
	if err != nil {
		return nil, err
	}

	psw, err = getUserActivePassword(user.Id)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, psw.Hash) {
		return nil, errors.New("email or password doesn't match")
	}

	T, err = authentication.NewAccessToken(user)
	if err != nil {
		return nil, err
	}

	bearer = &model.BearerToken{
		T: T,
	}

	return bearer, err
}
