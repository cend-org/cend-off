package profile

import (
	"context"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/configuration"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/media"
)

const (
	UserProfileImage = 3
)

func GetProfileImage(ctx context.Context) (*string, error) {
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

	err = database.Get(&media,
		`SELECT media.*
			FROM media
					 JOIN user_media_detail ON media.xid = user_media_detail.document_xid
					 JOIN user ON user.id = user_media_detail.owner_id
			WHERE user_media_detail.owner_id = ? AND user_media_detail.document_type = ?`, tok.UserId, UserProfileImage)
	if err != nil {
		return &networkLink, errx.DbGetError
	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/api/public/" + media.Xid + media.Extension

	return &networkLink, nil
}

func GetProfileImageThumb(ctx context.Context) (*string, error) {
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

	media, err = mediafile.GetMediaThumb(tok.UserId, UserProfileImage)
	if err != nil {
		return &networkLink, errx.DbGetError

	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/api/public/" + media.Xid + media.Extension

	return &networkLink, nil
}

func RemoveProfileImage(ctx context.Context) (*bool, error) {
	var (
		media           model.Media
		tok             *token.Token
		err             error
		status          bool
		userMediaDetail model.UserMediaDetail
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &status, errx.UnAuthorizedError
	}
	media, err = mediafile.GetMedia(tok.UserId, UserProfileImage)
	if err != nil {
		return &status, errx.DbGetError
	}

	mediaThumb, err := mediafile.GetMediaThumb(tok.UserId, UserProfileImage)
	if err != nil {
		return &status, errx.DbGetError
	}

	userMediaDetail, err = mediafile.GetUserMediaDetail(tok.UserId, UserProfileImage)
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
