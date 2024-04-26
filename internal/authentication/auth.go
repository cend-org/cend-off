package authentication

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"time"
)

func NewAccessToken(usr model.User) (string, error) {
	var (
		auth   model.Authorization
		tok    token.Token
		access string
		err    error
	)

	err = database.Get(&auth, `SELECT * FROM authorization WHERE user_id = ? ORDER BY level desc limit 1`, usr.ID)
	if err != nil {
		return "", err
	}

	/* fill token field */
	tok.UserID = usr.ID
	tok.UserLevel = auth.Level

	tok.ExpirationDate.Value = time.Now().Add(time.Hour * 24)

	/* parse token */
	access, err = token.New(tok)
	if err != nil {
		return access, err
	}

	time.Sleep(time.Second * 2)

	return access, err
}
