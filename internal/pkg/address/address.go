package address

import (
	"context"
	"duval/internal/authentication"
	"duval/internal/graph/model"
	"duval/internal/utils/errx"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"github.com/joinverse/xid"
)

func NewAddress(ctx *context.Context, input *model.NewAddress) (*model.Address, error) {

	var (
		tok         *authentication.Token
		isUser      model.UserAddress
		userId      uint
		address     model.Address
		userAddress model.UserAddress
		err         error
	)
	address.Country = input.Country
	address.City = input.City
	address.Latitude = input.Latitude
	address.Longitude = input.Longitude
	address.Street = input.Street
	address.FullAddress = input.FullAddress

	tok, err = authentication.GetTokenDataFromContext(*ctx)
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

/*
UPDATE ADDRESS OF A USER BY PROVIDING ID IN THE BODY
*/
func UpdateUserAddress(ctx *context.Context, input *model.NewAddress) (*model.Address, error) {
	var (
		address     model.Address
		userAddress model.UserAddress
		err         error
		tok         *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &address, errx.UnAuthorizedError
	}

	userAddress, err = GetUserAddressWithId(uint(tok.UserId))
	if userAddress.AddressId == state.ZERO {
		return &address, errx.UnAuthorizedError
	}

	address.Country = input.Country
	address.City = input.City
	address.Latitude = input.Latitude
	address.Longitude = input.Longitude
	address.Street = input.Street
	address.FullAddress = input.FullAddress

	err = database.Update(address)
	if err != nil {
		return &address, errx.DbUpdateError
	}

	return &address, nil
}

/*

	GET USER ADDRESS  BASED ON user_id PROVIDED IN PARAMS

*/

func GetUserAddress(ctx *context.Context) (*model.Address, error) {
	var (
		tok *authentication.Token

		userId  uint
		address model.Address
		err     error
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &address, errx.UnAuthorizedError
	}

	userId = uint(tok.UserId)

	err = database.Get(&address, `SELECT address.*
   FROM address JOIN user_address
   ON address.id = user_address.address_id
   WHERE user_address.user_id = ?`, userId)
	if err != nil {
		return &address, errx.DbGetError
	}

	return &address, nil
}

/*

	REMOVE USER ADDRESS  BASED ON user_id PROVIDED IN PARAMS

*/

func RemoveUserAddress(ctx *context.Context) (string, error) {
	var (
		tok *authentication.Token

		userId      uint
		address     model.Address
		userAddress model.UserAddress
		err         error
		status      string
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
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

	GET USER_ADDRESS WITH USER_ID

*/

func GetUserAddressWithId(userId uint) (userAddress model.UserAddress, err error) {
	err = database.Get(&userAddress, "SELECT * FROM user_address Where user_id = ?", userId)
	if err != nil {
		return userAddress, err
	}
	return userAddress, err
}
