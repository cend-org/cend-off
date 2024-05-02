package post

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
)

func NewPost(ctx context.Context, input *model.PostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: NewPost - newPost"))
}

func UpdPost(ctx context.Context, input *model.PostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: UpdPost - updPost"))
}

func RemovePost(ctx context.Context, postId int) (*string, error) {
	panic(fmt.Errorf("not implemented: RemovePost - removePost"))
}

func TagPost(ctx context.Context, input *model.PostTagInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: TagPost - tagPost"))
}

func UpdTagOnPost(ctx context.Context, input *model.PostTagInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: UpdTagOnPost - updTagOnPost"))
}

func RemoveTagOnPost(ctx context.Context, postId int) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: RemoveTagOnPost - removeTagOnPost"))
}

func GetPosts(ctx context.Context) ([]model.Post, error) {
	panic(fmt.Errorf("not implemented: GetPosts - getPosts"))
}

func ViewPost(ctx context.Context, postID int) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: ViewPost - viewPost"))
}

func GetTaggedPost(ctx context.Context, postId int) ([]model.Post, error) {
	panic(fmt.Errorf("not implemented: GetTaggedPost - getTaggedPost"))
}

func SearchPost(ctx context.Context, keyword string) ([]model.Post, error) {
	panic(fmt.Errorf("not implemented: SearchPost - searchPost"))
}
