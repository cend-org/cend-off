package authentication

import (
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func NewAccessToken(usr model.User) (string, error) {
	var (
		auth   model.Authorization
		tok    token.Token
		access string
		err    error
	)

	err = database.Get(&auth, `SELECT * FROM authorization WHERE user_id = ? ORDER BY level desc limit 1`, usr.Id)
	if err != nil {
		return "", err
	}

	/* fill token field */
	tok.UserId = usr.Id
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

func GinContext(ctx *gin.Context) (tok *token.Token, err error) {
	tokenString := ctx.GetHeader("Authorization")
	if len(strings.TrimSpace(tokenString)) == 0 {
		return nil, errors.New("bad header value given")
	}

	bearer := strings.Split(tokenString, " ")
	if len(bearer) != 2 {
		return nil, errors.New("incorrectly formatted authorization header")
	}

	return token.Parse(bearer[1]), err
}
