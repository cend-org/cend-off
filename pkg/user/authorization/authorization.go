package authorization

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
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
