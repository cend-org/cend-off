package phone

import (
	"context"
	"duval/internal/authentication"
	"duval/internal/graph/model"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/pkg/database"
	"errors"
)

func NewPhoneNumber(ctx *context.Context, input *model.NewPhoneNumber) (*model.PhoneNumber, error) {
	var (
		userPhoneNumber model.UserPhoneNumber
		newPhone        model.PhoneNumber
		err             error
		tok             *authentication.Token
	)
	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &newPhone, errx.UnAuthorizedError
	}

	if !utils.IsValidPhone(newPhone.MobilePhoneNumber) {
		return &newPhone, errx.ParseError
	}
	newPhone.MobilePhoneNumber = input.MobilePhoneNumber
	newPhone.IsUrgency = input.IsUrgency
	newPhone.Id, err = database.InsertOne(newPhone)
	if err != nil {
		return &newPhone, errx.Lambda(err)
	}
	// Link phone to user.
	userPhoneNumber.UserId = tok.UserId
	userPhoneNumber.PhoneNumberId = newPhone.Id
	_, err = database.InsertOne(userPhoneNumber)
	if err != nil {
		return &newPhone, errx.DbInsertError
	}

	return &newPhone, nil
}

func UpdateUserPhoneNumber(ctx *context.Context, input *model.NewPhoneNumber) (*model.PhoneNumber, error) {
	var (
		phoneNumber  model.PhoneNumber
		currentPhone model.PhoneNumber
		err          error
		tok          *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &phoneNumber, errx.UnAuthorizedError
	}

	currentPhone, err = GetPhoneById(int(tok.UserId))
	if currentPhone.Id == 0 {
		return &phoneNumber, errx.Lambda(errors.New("create new phone number instead"))
	}

	if !utils.IsValidPhone(phoneNumber.MobilePhoneNumber) {
		return &phoneNumber, errx.ParseError
	}

	err = database.Update(phoneNumber)
	if err != nil {
		return &phoneNumber, errx.DbUpdateError
	}

	return &phoneNumber, nil
}

func GetUserPhoneNumber(ctx *context.Context) (*model.PhoneNumber, error) {
	var (
		phone model.PhoneNumber
		err   error
		tok   *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(*ctx)
	if err != nil {
		return &phone, errx.UnAuthorizedError
	}

	err = database.Get(&phone, `SELECT phone_number.mobile_phone_number
	FROM phone_number JOIN user_phone_number 
	ON phone_number.id = user_phone_number.id 
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
	err = database.Get(&phoneNumber, `SELECT * FROM phone_number JOIN user_phone_number ON phone_number.id = user_phone_number.phone_number_id WHERE user_phone_number.user_id = ?`, userId)
	if err != nil {
		return phoneNumber, err
	}

	return phoneNumber, nil
}
