package user

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
)

func UpdateProfileAndPassword(userId int, new model.UserInput, newPass model.PasswordInput) (usr *model.User, err error) {
	var user model.User

	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, userId)
	if err != nil {
		return nil, err
	}

	user = model.MapUserInputToUser(new, user)

	_, err = NewPassword(user.Id, newPass)
	if err != nil {
		return nil, err
	}

	err = database.Update(user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func UpdMyProfile(userId int, new model.UserInput) (usr *model.User, err error) {
	var user model.User

	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, userId)
	if err != nil {
		return nil, err
	}

	user = model.MapUserInputToUser(new, user)

	err = database.Update(user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func MyProfile(userId int) (usr *model.User, err error) {
	var user model.User
	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, userId)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func getUserByEmail(email string) (user model.User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return user, err
	}
	return user, err
}
