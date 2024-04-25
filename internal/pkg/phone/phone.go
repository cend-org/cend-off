package phone

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
)

func NewPhoneNumber(ctx context.Context, input *model.PhoneNumberInput) (*model.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented: NewPhoneNumber - newPhoneNumber"))
}

func UpdateUserPhoneNumber(ctx context.Context, input *model.PhoneNumberInput) (*model.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented: UpdateUserPhoneNumber - updateUserPhoneNumber"))
}

func GetUserPhoneNumber(ctx context.Context) (*model.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented: GetUserPhoneNumber - getUserPhoneNumber"))
}
