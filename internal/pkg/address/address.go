package address

import (
	"context"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/joinverse/xid"
	"time"
)

func NewAddress(ctx context.Context, input *model.AddressInput) (*model.Address, error) {

	var (
		tok         *token.Token
		isUser      model.UserAddress
		userId      int
		address     model.Address
		userAddress model.UserAddress
		err         error
	)
	address = model.MapAddressInputToAddress(*input, address)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &address, errx.UnAuthorizedError
	}
	userId = tok.UserId

	// get user address
	isUser, err = GetUserAddressWithId(userId)
	if isUser.AddressId > state.ZERO {
		return &address, errx.DuplicateAddressError
	}
	address.Xid = xid.New().String()

	address.Id, err = database.InsertOne(address)
	if err != nil {
		return &address, errx.DbInsertError
	}

	// Link new address to the current user
	userAddress.UserId = userId
	userAddress.AddressId = address.Id
	_, err = database.InsertOne(userAddress)
	if err != nil {
		return &address, errx.LinkUserError
	}

	return &address, nil
}

func UpdateUserAddress(ctx context.Context, input *model.AddressInput) (*model.Address, error) {
	var (
		address     model.Address
		userAddress model.UserAddress
		err         error
		tok         *token.Token
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &address, errx.UnAuthorizedError
	}

	userAddress, err = GetUserAddressWithId(int(tok.UserId))
	if userAddress.AddressId == state.ZERO {
		return &address, errx.UnAuthorizedError
	}

	address = model.MapAddressInputToAddress(*input, address)

	err = database.Update(address)
	if err != nil {
		return &address, errx.DbUpdateError
	}

	return &address, nil
}

func GetUserAddress(ctx context.Context) (*model.Address, error) {
	var (
		tok *token.Token

		userId  int
		address model.Address
		err     error
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &address, errx.UnAuthorizedError
	}

	userId = int(tok.UserId)

	err = database.Get(&address, `SELECT address.*
   FROM address JOIN user_address
   ON address.id = user_address.address_id
   WHERE user_address.user_id = ?`, userId)
	if err != nil {
		return &address, errx.DbGetError
	}

	return &address, nil
}

func RemoveUserAddress(ctx context.Context) (string, error) {
	var (
		tok *token.Token

		userId      int
		address     model.Address
		userAddress model.UserAddress
		err         error
		status      string
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return status, errx.UnAuthorizedError
	}
	userId = tok.UserId

	err = database.Get(&address, `SELECT address.*
   FROM address JOIN user_address
   ON address.id = user_address.address_id
   WHERE user_address.user_id = ?`, userId)
	if err != nil {
		return status, errx.DbGetError
	}
	time.Sleep(100)
	err = database.Delete(address)
	if err != nil {
		return status, errx.DbDeleteError
	}
	// and remove user_address

	err = database.Get(&userAddress, `SELECT * FROM user_address where user_id = ?`, userId)
	if err != nil {
		return status, errx.DbGetError
	}

	err = database.Delete(userAddress)
	if err != nil {
		return status, errx.DbDeleteError
	}
	status = "Address removed successfully!"

	return status, nil
}

/*
	UTILS
*/

func GetUserAddressWithId(userId int) (userAddress model.UserAddress, err error) {
	err = database.Get(&userAddress, "SELECT * FROM user_address Where user_id = ?", userId)
	if err != nil {
		return userAddress, err
	}
	return userAddress, err
}
