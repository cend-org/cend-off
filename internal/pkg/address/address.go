package address

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
)

func NewAddress(ctx context.Context, input model.AddressInput) (*model.Address, error) {

	panic(fmt.Errorf("not implemented: NewAddress - newAddress"))

}

func UpdateUserAddress(ctx context.Context, input *model.AddressInput) (*model.Address, error) {
	panic(fmt.Errorf("not implemented: UpdateUserAddress - updateUserAddress"))
}

// getUserAddress: Address!
func GetUserAddress(ctx context.Context) (*model.Address, error) {
	panic(fmt.Errorf("not implemented: GetUserAddress - getUserAddress"))
}

//removeUserAddress: String!

func RemoveUserAddress(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: RemoveUserAddress - removeUserAddress"))
}
