package cv

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/configuration"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/media"
)

const (
	CV = 0
)

func RemoveProfileCv(userId int) (bool, error) {
	var (
		media     model.Media
		err       error
		status    bool
		filePath  string
		thumbPath string
	)

	if userId == state.ZERO {
		return status, errx.UnAuthorizedError
	}
	media, err = mediafile.GetMedia(userId, CV)
	if err != nil {
		return status, errx.DbGetError
	}

	mediaThumb, err := mediafile.GetMediaThumb(userId, CV)
	if err != nil {
		return status, errx.DbGetError
	}

	userMediaDetail, err := mediafile.GetUserMediaDetail(userId, CV)
	if err != nil {
		return status, errx.DbGetError
	}

	filePath = utils.FILE_UPLOAD_DIR + media.Xid + media.Extension
	thumbPath = utils.FILE_UPLOAD_DIR + utils.THUMB_FILE_UPLOAD_DIR + mediaThumb.Xid + mediaThumb.Extension

	err = mediafile.ClearMediaFile(filePath)
	if err != nil {
		return status, errx.SupportError
	}

	err = mediafile.ClearMediaFile(thumbPath)
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

func GetProfileCv(userId int) (string, error) {
	var (
		media       model.Media
		err         error
		networkLink string
	)

	err = database.Get(&media,
		`SELECT media.*
			FROM media
					 JOIN user_media_detail ON media.xid = user_media_detail.document_xid
					 JOIN user ON user.id = user_media_detail.owner_id
			WHERE user_media_detail.owner_id = ? AND user_media_detail.document_type = ?`, userId, CV)
	if err != nil {
		return networkLink, err
	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/public/" + media.Xid + media.Extension

	return networkLink, nil
}

func GetProfileCvThumb(userId int) (string, error) {
	var (
		media       model.MediaThumb
		err         error
		networkLink string
	)

	media, err = mediafile.GetMediaThumb(userId, CV)
	if err != nil {
		return networkLink, err

	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/public/thumb/" + media.Xid + media.Extension

	return networkLink, nil
}
