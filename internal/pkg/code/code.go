package code

import (
	"duval/internal/graph/model"
	"duval/pkg/database"
	"math/rand"
)

func NewUserVerificationCode(userId uint) (err error) {
	var code model.Code
	code.VerificationCode = rand.Intn(9999)
	code.UserId = userId

	_, err = database.InsertOne(code)
	if err != nil {
		return err
	}

	return err
}

func IsUserVerificationCodeValid(userId uint, verificationCode int) (err error) {
	var code model.Code
	var query = `SELECT * FROM code WHERE user_id = ? AND verification_code = ? ORDER BY created_at desc LIMIT 1`
	err = database.Get(&code, query, userId, verificationCode)
	if err != nil {
		return err
	}

	return err
}
