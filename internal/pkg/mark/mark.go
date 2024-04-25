package mark

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
)

func RateUser(ctx context.Context, input *model.MarkInput) (*model.Mark, error) {
	panic(fmt.Errorf("not implemented: RateUser - rateUser"))
}

func GetUserAverageMark(ctx context.Context, userId int) (*int, error) {
	panic(fmt.Errorf("not implemented: GetUserAverageMark - getUserAverageMark"))
}

func GetUserMarkComment(ctx context.Context) ([]model.Mark, error) {
	panic(fmt.Errorf("not implemented: GetUserMarkComment - getUserMarkComment"))
}
