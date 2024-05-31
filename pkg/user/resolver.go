package user

import (
	"context"
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/token"
)

type UserQuery struct{}
type UserMutation struct{}

// MyProfile is the resolver for the MyProfile field.
func (r *UserQuery) MyProfile(ctx context.Context) (*model.User, error) {
	var tok *token.Token
	var err error

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errors.New("unAuthorized")
	}

	return MyProfile(tok.UserId)
}
