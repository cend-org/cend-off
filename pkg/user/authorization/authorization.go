package authorization

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils/state"
)

const (
	STUDENT = 0
	PARENT  = 1
	TUTOR   = 2
	PROF    = 3
)

func NewUserAuthorization(userId, authorizationLevel int) (err error) {
	var (
		auth model.Authorization
	)

	auth.UserId = userId
	auth.Level = authorizationLevel

	_, err = database.InsertOne(auth)
	if err != nil {
		return err
	}
	return err
}

func GetUserAuthorization(userId, level int) (auth model.Authorization, err error) {
	err = database.Get(&auth, `SELECT * FROM authorization WHERE user_id = ? AND level = ?`, userId, level)
	if err != nil {
		return auth, err
	}

	return auth, err
}

func IsUserStudent(userId int) (ret bool) {
	return isUserHasAuthorizationLevel(userId, STUDENT)
}

func IsUserParent(userId int) (ret bool) {
	return isUserHasAuthorizationLevel(userId, PARENT)
}

func IsUserTutor(userId int) (ret bool) {
	return isUserHasAuthorizationLevel(userId, TUTOR)
}

func IsUserProfessor(userId int) (ret bool) {
	return isUserHasAuthorizationLevel(userId, PROF)
}

func isUserHasAuthorizationLevel(userId, authorizationLevel int) (ret bool) {
	var (
		err  error
		auth model.Authorization
	)

	auth, err = GetUserAuthorization(userId, authorizationLevel)
	if err != nil {
		return false
	}

	return auth.Id > state.ZERO
}
