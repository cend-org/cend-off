package post

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
)

func NewPost(ctx context.Context, input *model.PostInput) (*model.Post, error) {
	var (
		post model.Post
		err  error
		tok  *token.Token
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &post, errx.UnAuthorizedError
	}

	post = model.MapPostInputToPost(*input, post)

	post.PublisherId = tok.UserId

	_, err = database.InsertOne(post)
	if err != nil {
		return &post, errx.DbInsertError
	}
	return &post, nil
}

func UpdPost(ctx context.Context, input *model.PostInput, postId int) (*model.Post, error) {
	var (
		post model.Post
		err  error
		tok  *token.Token
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &post, errx.UnAuthorizedError
	}

	post, err = GetSingleUserPost(postId, tok.UserId)
	if err != nil {
		return &post, errx.DbGetError
	}

	post = model.MapPostInputToPost(*input, post)

	err = database.Update(post)
	if err != nil {
		return &post, errx.DbUpdateError
	}

	return &post, nil
}

func RemovePost(ctx context.Context, postId int) (*string, error) {
	var (
		post   model.Post
		err    error
		tok    *token.Token
		status string
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	post, err = GetSingleUserPost(postId, tok.UserId)
	if err != nil {
		return &status, errx.DbGetError
	}

	err = database.Delete(post)
	if err != nil {
		return &status, errx.DbInsertError
	}

	return &status, nil
}

func TagPost(ctx context.Context, input *model.PostTagInput) (*model.Post, error) {
	var (
		post    model.Post
		postTag model.PostTag
		err     error
		tok     *token.Token
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &post, errx.UnAuthorizedError
	}

	postTag = model.MapPostTagInputToPostTag(*input, postTag)

	post, err = GetSingleUserPost(postTag.PostId, tok.UserId)
	if err != nil {
		return &post, errx.DbGetError
	}

	_, err = database.InsertOne(postTag)
	if err != nil {
		return &post, errx.DbInsertError
	}

	return &post, nil
}

func UpdTagOnPost(ctx context.Context, input *model.PostTagInput) (*model.Post, error) {
	var (
		post    model.Post
		postTag model.PostTag
		err     error
		tok     *token.Token
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &post, errx.UnAuthorizedError
	}

	postTag, err = GetPostSinglePostTag(*input.PostId, tok.UserId)
	if err != nil {
		return &post, errx.DbGetError
	}

	postTag = model.MapPostTagInputToPostTag(*input, postTag)

	post, err = GetSingleUserPost(postTag.PostId, tok.UserId)
	if err != nil {
		return &post, errx.DbGetError
	}

	err = database.Update(postTag)
	if err != nil {
		return &post, errx.DbInsertError
	}

	return &post, nil
}

func RemoveTagOnPost(ctx context.Context, postId int) (*model.Post, error) {
	var (
		post    model.Post
		postTag model.PostTag
		err     error
		tok     *token.Token
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &post, errx.UnAuthorizedError
	}

	postTag, err = GetPostSinglePostTag(postId, tok.UserId)
	if err != nil {
		return &post, errx.DbGetError
	}

	post, err = GetSingleUserPost(postTag.PostId, tok.UserId)
	if err != nil {
		return &post, errx.DbGetError
	}

	err = database.Delete(postTag)
	if err != nil {
		return &post, errx.DbInsertError
	}

	return &post, nil
}

func GetPosts(ctx context.Context) ([]model.Post, error) {
	var (
		posts []model.Post
		err   error
	)

	err = database.GetMany(&posts, `SELECT post.* FROM post LIMIT 10`)
	if err != nil {
		return posts, errx.DbGetError
	}

	return posts, nil
}

func ViewPost(ctx context.Context, postId int) (*model.Post, error) {
	var (
		post     model.Post
		postView model.PostView
		err      error
		tok      *token.Token
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &post, errx.UnAuthorizedError
	}

	post, err = GetSinglePost(postId)
	if err != nil {
		return &post, errx.DbGetError
	}

	postView.PostId = postId
	postView.ViewerId = tok.UserId

	_, err = database.InsertOne(postView)
	if err != nil {
		return &post, errx.DbInsertError
	}

	return &post, nil
}

func SearchPost(ctx context.Context, keyword string) ([]model.Post, error) {
	panic(fmt.Errorf("not implemented: SearchPost - searchPost"))
}

/*
	UTILS
*/

func GetSinglePost(postId int) (post model.Post, err error) {
	err = database.Get(&post, `SELECT post.* FROM post WHERE post.id = ? `, postId)
	if err != nil {
		return post, err
	}
	return post, nil
}

func GetSingleUserPost(postId, publisherId int) (post model.Post, err error) {
	err = database.Get(&post, `SELECT post.* FROM post WHERE post.id = ? AND post.publisher_id = ? `, postId, publisherId)
	if err != nil {
		return post, err
	}
	return post, nil
}

func GetPostSinglePostTag(postId, userId int) (postTag model.PostTag, err error) {
	err = database.Get(&postTag,
		`SELECT post_tag.*
			FROM post_tag
					 JOIN post ON post_tag.post_id = post.id
			WHERE post.id = ?
			  AND post.publisher_id = ?`, postId, userId)
	if err != nil {
		return postTag, err
	}
	return postTag, nil
}
