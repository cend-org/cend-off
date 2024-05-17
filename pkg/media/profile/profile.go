package profile

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/configuration"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/media"
)

const (
	UserProfileImage = 3
)

func GetProfileImage(userId int) (string, error) {
	var (
		media       model.Media
		err         error
		networkLink string
	)

	media, err = mediafile.GetMedia(userId, UserProfileImage)
	if err != nil {
		return networkLink, errx.DbGetError
	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/public/" + media.Xid + media.Extension

	return networkLink, nil
}

func GetProfileImageThumb(userId int) (string, error) {
	var (
		media       model.MediaThumb
		err         error
		networkLink string
	)

	media, err = mediafile.GetMediaThumb(userId, UserProfileImage)
	if err != nil {
		return networkLink, errx.DbGetError

	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/public/thumb/" + media.Xid + media.Extension

	return networkLink, nil
}

func RemoveProfileImage(userId int) (bool, error) {
	var (
		media           model.Media
		err             error
		status          bool
		userMediaDetail model.UserMediaDetail
	)

	if userId == state.ZERO {
		return status, errx.UnAuthorizedError
	}

	media, err = mediafile.GetMedia(userId, UserProfileImage)
	if err != nil {
		return status, errx.DbGetError
	}

	mediaThumb, err := mediafile.GetMediaThumb(userId, UserProfileImage)
	if err != nil {
		return status, errx.DbGetError
	}

	userMediaDetail, err = mediafile.GetUserMediaDetail(userId, UserProfileImage)
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
