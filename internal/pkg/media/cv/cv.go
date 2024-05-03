package cv

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/configuration"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/pkg/media"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/gabriel-vasile/mimetype"
	"github.com/joinverse/xid"
	"path/filepath"
)

const (
	CV = 0
)

func UploadProfileLetter(ctx context.Context, file *graphql.Upload) (*model.Media, error) {
	var (
		media        model.Media
		tok          *token.Token
		err          error
		documentType int
	)
	documentType = CV

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &media, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &media, errx.UnAuthorizedError
	}

	mType, err := mimetype.DetectReader(file.File)
	if err != nil {
		return &media, errx.TypeError

	}
	if !utils.IsValidFile(mType.String()) {
		return &media, errx.TypeError
	}

	media.FileName = file.Filename
	media.Extension = filepath.Ext(file.Filename)
	media.Xid = xid.New().String()
	uploadPath := "./" + utils.FILE_UPLOAD_DIR + media.Xid + media.Extension

	err = mediafile.SaveFile(uploadPath, *file)
	if err != nil {
		return &media, errx.Lambda(err)
	}

	_, err = database.InsertOne(media)
	if err != nil {
		return &media, errx.DbInsertError
	}

	err = utils.CreateDocumentThumb(media.Xid, media.Extension, *file)
	if err != nil {
		return &media, errx.ThumbError
	}

	err = mediafile.SetUserMediaDetail(documentType, tok.UserId, media.Xid)
	if err != nil {
		return &media, errx.DbInsertError
	}

	return &media, nil
}

func UpdateProfileCv(ctx context.Context, file *graphql.Upload) (*model.Media, error) {
	var (
		media model.Media
		tok   *token.Token
		err   error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &media, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &media, errx.UnAuthorizedError

	}

	mType, err := mimetype.DetectReader(file.File)
	if err != nil {
		return &media, errx.TypeError

	}

	if !utils.IsValidDocument(mType.String()) {
		return &media, errx.TypeError
	}

	oldMedia, err := mediafile.GetMedia(tok.UserId, CV)
	if err != nil {
		return &media, errx.DbGetError
	}

	media.FileName = file.Filename
	media.Extension = filepath.Ext(file.Filename)
	media.Xid = oldMedia.Xid

	err = mediafile.RemoveMedia(oldMedia)
	if err != nil {
		return &media, errx.DbDeleteError
	}

	uploadPath := "./" + utils.FILE_UPLOAD_DIR + media.Xid + media.Extension

	err = mediafile.SaveFile(uploadPath, *file)
	if err != nil {
		return &media, errx.Lambda(err)
	}

	_, err = database.InsertOne(media)
	if err != nil {
		return &media, errx.DbInsertError
	}

	err = utils.CreateDocumentThumb(media.Xid, media.Extension, *file)
	if err != nil {
		return &media, errx.ThumbError
	}

	return &media, nil
}

func RemoveProfileCv(ctx context.Context) (*string, error) {
	var (
		media  model.Media
		tok    *token.Token
		err    error
		status string
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &status, errx.UnAuthorizedError
	}
	media, err = mediafile.GetMedia(tok.UserId, CV)
	if err != nil {
		return &status, errx.DbGetError
	}

	err = mediafile.RemoveMedia(media)
	if err != nil {
		return &status, errx.DbDeleteError
	}

	userMediaDetail, err := mediafile.GetUserMediaDetail(tok.UserId, CV)
	if err != nil {
		return &status, errx.DbGetError
	}

	err = mediafile.RemoveUserMediaDetail(userMediaDetail)
	if err != nil {
		return &status, errx.DbDeleteError
	}
	status = "success"
	return &status, nil

}

func GetProfileCv(ctx context.Context) (*string, error) {
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
			WHERE user_media_detail.owner_id = ? AND user_media_detail.document_type = ?`, tok.UserId, CV)
	if err != nil {
		return &networkLink, errx.DbGetError
	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/api/public/" + media.Xid + media.Extension

	return &networkLink, nil
}

func GetProfileCvThumb(ctx context.Context) (*string, error) {
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

	media, err = mediafile.GetMediaThumb(tok.UserId, CV)
	if err != nil {
		return &networkLink, errx.DbGetError

	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/api/public/" + media.Xid + media.Extension

	return &networkLink, nil
}
