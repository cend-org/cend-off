package video

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/cend-org/duval/graph/model"
)

func UpdateProfileVideo(ctx context.Context, file *graphql.Upload) (*model.Media, error) {
	panic(fmt.Errorf("not implemented: UpdateProfileVideo - updateProfileVideo"))
}

func RemoveProfileVideo(ctx context.Context, mediaID int) (*string, error) {
	panic(fmt.Errorf("not implemented: RemoveProfileVideo - removeProfileVideo"))
}

func GetProfileVideo(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: GetProfileVideo - getProfileVideo"))
}

func GetProfileVideoThumb(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: GetProfileVideoThumb - getProfileVideoThumb"))
}
