package cover

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/configuration"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/media"
)

const (
	Letter = 1
)

func GetProfileLetter(userId int) (string, error) {
	var (
		media       model.Media
		err         error
		networkLink string
	)

	media, err = mediafile.GetMedia(userId, Letter)
	if err != nil {
		return networkLink, errx.DbGetError
	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/public/" + media.Xid + media.Extension

	return networkLink, nil
}

func GetProfileLetterThumb(userId int) (string, error) {

	var (
		media       model.MediaThumb
		err         error
		networkLink string
	)

	media, err = mediafile.GetMediaThumb(userId, Letter)
	if err != nil {
		return networkLink, errx.DbGetError

	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/public/thumb/" + media.Xid + media.Extension

	return networkLink, nil
}

func RemoveProfileLetter(userId int) (bool, error) {
	var (
		media  model.Media
		err    error
		status bool
	)

	if userId == state.ZERO {
		return status, errx.UnAuthorizedError
	}
	media, err = mediafile.GetMedia(userId, Letter)
	if err != nil {
		return status, errx.DbGetError
	}

	mediaThumb, err := mediafile.GetMediaThumb(userId, Letter)
	if err != nil {
		return status, errx.DbGetError
	}

	userMediaDetail, err := mediafile.GetUserMediaDetail(userId, Letter)
	if err != nil {
		return status, errx.DbGetError
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
