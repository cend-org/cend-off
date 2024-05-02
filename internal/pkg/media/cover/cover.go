package cover

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/cend-org/duval/graph/model"
)

func GetProfileLetter(ctx context.Context) (*model.Media, error) {
	panic(fmt.Errorf("not implemented: GetProfileLetter - getProfileLetter"))
}

func GetProfileLetterThumb(ctx context.Context) (*model.MediaThumb, error) {
	panic(fmt.Errorf("not implemented: GetProfileLetterThumb - getProfileLetterThumb"))
}

func UpdateProfileLetter(ctx context.Context, file *graphql.Upload) (*model.Media, error) {
	panic(fmt.Errorf("not implemented: UpdateProfileLetter - updateProfileLetter"))
}

func RemoveProfileLetter(ctx context.Context, mediaID int) (*string, error) {
	panic(fmt.Errorf("not implemented: RemoveProfileLetter - removeProfileLetter"))
}
