package code

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
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
