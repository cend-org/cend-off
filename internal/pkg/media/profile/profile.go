package profile

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/cend-org/duval/graph/model"
)

func GetProfileImage(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: GetProfileImage - getProfileImage"))
}

func GetProfileImageThumb(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: GetProfileImageThumb - getProfileImageThumb"))
}

func UpdateProfileImage(ctx context.Context, file *graphql.Upload) (*model.Media, error) {
	panic(fmt.Errorf("not implemented: UpdateProfileImage - updateProfileImage"))
}

func RemoveProfileImage(ctx context.Context, mediaID int) (*string, error) {
	panic(fmt.Errorf("not implemented: RemoveProfileImage - removeProfileImage"))
}
