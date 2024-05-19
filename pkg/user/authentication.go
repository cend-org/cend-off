package user

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/authentication"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
)

func Login(email string, password string) (bearer *model.BearerToken, err error) {
	var user model.User
	var psw model.Password
	var T string

	user, err = getUserByEmail(email)
	if err != nil {
		return nil, errx.ToRegisterEmailError
	}

	psw, err = getUserActivePassword(user.Id)
	if err != nil {
		return nil, errx.StatusNeedPasswordError
	}

	if !utils.CheckPasswordHash(password, psw.Hash) {
		return nil, errx.IncorrectPasswordError
	}

	T, err = authentication.NewAccessToken(user)
	if err != nil {
		return nil, errx.SupportError
	}

	bearer = &model.BearerToken{
		T: T,
	}

	return bearer, err
}
