package phone

import (
	"context"
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
)

func NewPhoneNumber(ctx context.Context, input *model.PhoneNumberInput) (*model.PhoneNumber, error) {
	var (
		userPhoneNumber model.UserPhoneNumber
		phone           model.PhoneNumber
		err             error
		tok             *token.Token
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &phone, errx.UnAuthorizedError
	}

	if !utils.IsValidPhone(*input.MobilePhoneNumber) {
		return &phone, errx.ParseError
	}
	phone = model.MapPhoneNumberInputToPhoneNumber(*input, phone)

	phone.Id, err = database.InsertOne(phone)
	if err != nil {
		return &phone, errx.Lambda(err)
	}

	// Link phone to user.
	userPhoneNumber.UserId = tok.UserId
	userPhoneNumber.PhoneNumberId = phone.Id

	_, err = database.InsertOne(userPhoneNumber)
	if err != nil {
		return &phone, errx.DbInsertError
	}

	return &phone, nil
}

func UpdateUserPhoneNumber(ctx context.Context, input *model.PhoneNumberInput) (*model.PhoneNumber, error) {
	var (
		phoneNumber model.PhoneNumber
		err         error
		tok         *token.Token
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &phoneNumber, errx.UnAuthorizedError
	}

	phoneNumber, err = GetPhoneById(tok.UserId)
	if phoneNumber.Id == 0 {
		return &phoneNumber, errx.Lambda(errors.New("create new phone number instead"))
	}

	if !utils.IsValidPhone(*input.MobilePhoneNumber) {
		return &phoneNumber, errx.ParseError
	}
	phoneNumber = model.MapPhoneNumberInputToPhoneNumber(*input, phoneNumber)

	err = database.Update(phoneNumber)
	if err != nil {
		return &phoneNumber, errx.DbUpdateError
	}

	return &phoneNumber, nil
}

func GetUserPhoneNumber(ctx context.Context) (*model.PhoneNumber, error) {
	var (
		phone model.PhoneNumber
		err   error
		tok   *token.Token
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &phone, errx.UnAuthorizedError
	}

	err = database.Get(&phone,
		`SELECT phone_number.* 
				FROM phone_number 
    			JOIN user_phone_number ON phone_number.id = user_phone_number.phone_number_id 
				WHERE user_phone_number.user_id = ?`, tok.UserId)
	if err != nil {
		return &phone, errx.DbGetError
	}

	return &phone, nil
}

/*

	UTILS

*/

func GetPhoneById(userId int) (phoneNumber model.PhoneNumber, err error) {
	err = database.Get(&phoneNumber, `SELECT phone_number.* FROM phone_number JOIN user_phone_number ON phone_number.id = user_phone_number.phone_number_id WHERE user_phone_number.user_id = ?`, userId)
	if err != nil {
		return phoneNumber, err
	}

	return phoneNumber, nil
}
