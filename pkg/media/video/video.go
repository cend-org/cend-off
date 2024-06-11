package video

import (
	"context"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/configuration"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/media"
)

const (
	UserProfileVideo = 2
)

func GetProfileVideo(userId int) (string, error) {
	var (
		media       model.Media
		err         error
		networkLink string
	)

	media, err = mediafile.GetMedia(userId, UserProfileVideo)
	if err != nil {
		return networkLink, errx.DbGetError
	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/public/" + media.Xid + media.Extension

	return networkLink, nil
}

func GetProfileVideoThumb(ctx context.Context) (*string, error) {
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

	media, err = mediafile.GetMediaThumb(tok.UserId, UserProfileVideo)
	if err != nil {
		return &networkLink, errx.DbGetError

	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/public/thumb/" + media.Xid + media.Extension

	return &networkLink, nil
}

func RemoveProfileVideo(userId int) (bool, error) {
	var (
		media           model.Media
		err             error
		status          bool
		userMediaDetail model.UserMediaDetail
		filePath        string
	)

	if userId == state.ZERO {
		return status, errx.UnAuthorizedError
	}
	media, err = mediafile.GetMedia(userId, UserProfileVideo)
	if err != nil {
		return status, errx.DbGetError
	}

	mediaThumb, err := mediafile.GetMediaThumb(userId, UserProfileVideo)
	if err != nil {
		return status, errx.DbGetError
	}

	userMediaDetail, err = mediafile.GetUserMediaDetail(userId, UserProfileVideo)
	if err != nil {
		return status, errx.DbGetError
	}

	filePath = utils.FILE_UPLOAD_DIR + media.Xid + media.Extension

	err = mediafile.ClearMediaFile(filePath)
	if err != nil {
		return status, errx.SupportError
	}

	err = mediafile.RemoveMedia(media)
	if err != nil {
		return status, errx.DbDeleteError
	}

	err = mediafile.RemoveMediaThumb(mediaThumb)
	if err != nil {
		return status, errx.DbDeleteError
	}

	err = mediafile.RemoveUserMediaDetail(userMediaDetail)
	if err != nil {
		return status, errx.DbDeleteError
	}

	status = true

	return status, nil

}
