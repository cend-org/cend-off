package cv

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/cend-org/duval/graph/model"
)

func UpdateProfileCv(ctx context.Context, file *graphql.Upload) (*model.Media, error) {
	panic(fmt.Errorf("not implemented: UpdateProfileCv - updateProfileCv"))
}

func RemoveProfileCv(ctx context.Context, mediaID int) (*string, error) {
	panic(fmt.Errorf("not implemented: RemoveProfileCv - removeProfileCv"))
}

func GetProfileCv(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: GetProfileCv - getProfileCv"))
}

func GetProfileCvThumb(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: GetProfileCvThumb - getProfileCvThumb"))
}
