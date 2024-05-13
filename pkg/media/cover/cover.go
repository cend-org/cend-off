package cover

import (
	"context"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/configuration"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/media"
)

const (
	Letter = 1
)

func GetProfileLetter(ctx context.Context) (*string, error) {
	var (
		tok         *token.Token
		media       model.Media
		err         error
		networkLink string
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &networkLink, errx.UnAuthorizedError
	}

	media, err = mediafile.GetMedia(tok.UserId, Letter)
	if err != nil {
		return &networkLink, errx.DbGetError
	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/public/" + media.Xid + media.Extension

	return &networkLink, nil
}

func GetProfileLetterThumb(ctx context.Context) (*string, error) {

	var (
		tok         *token.Token
		media       model.MediaThumb
		err         error
		networkLink string
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &networkLink, errx.UnAuthorizedError
	}

	media, err = mediafile.GetMediaThumb(tok.UserId, Letter)
	if err != nil {
		return &networkLink, errx.DbGetError

	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/public/thumb/" + media.Xid + media.Extension

	return &networkLink, nil
}

func RemoveProfileLetter(ctx context.Context) (*bool, error) {
	var (
		media  model.Media
		tok    *token.Token
		err    error
		status bool
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &status, errx.UnAuthorizedError
	}
	media, err = mediafile.GetMedia(tok.UserId, Letter)
	if err != nil {
		return &status, errx.DbGetError
	}

	mediaThumb, err := mediafile.GetMediaThumb(tok.UserId, Letter)
	if err != nil {
		return &status, errx.DbGetError
	}

	userMediaDetail, err := mediafile.GetUserMediaDetail(tok.UserId, Letter)
	if err != nil {
		return &status, errx.DbGetError
	}

	err = mediafile.RemoveMedia(media)
	if err != nil {
		return &status, errx.DbDeleteError
	}

	err = mediafile.RemoveMediaThumb(mediaThumb)
	if err != nil {
		return &status, errx.DbDeleteError
	}

	err = mediafile.RemoveUserMediaDetail(userMediaDetail)
	if err != nil {
		return &status, errx.DbDeleteError
	}

	status = true

	return &status, nil

}
