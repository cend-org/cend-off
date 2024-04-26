package code

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"math/rand"
)

func GetCode(ctx context.Context) (*model.Code, error) {
	panic(fmt.Errorf("not implemented: GetCode - getCode"))
}

func VerifyUserEmailValidationCode(ctx context.Context, code int) (int, error) {
	panic(fmt.Errorf("not implemented: VerifyUserEmailValidationCode - verifyUserEmailValidationCode"))
}

func SendUserEmailValidationCode(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: SendUserEmailValidationCode - sendUserEmailValidationCode"))
}

/*
	UTILS
*/

func NewUserVerificationCode(userId int) (err error) {
	var code model.Code
	code.VerificationCode = rand.Intn(9999)
	code.UserID = userId

	_, err = database.InsertOne(code)
	if err != nil {
		return err
	}

	return err
}

func IsUserVerificationCodeValid(userId int, verificationCode int) (err error) {
	var code model.Code
	var query = `SELECT * FROM code WHERE user_id = ? AND verification_code = ? ORDER BY created_at desc LIMIT 1`
	err = database.Get(&code, query, userId, verificationCode)
	if err != nil {
		return err
	}

	return err
}
