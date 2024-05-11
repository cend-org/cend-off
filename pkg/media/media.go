package mediafile

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/authentication"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/joinverse/xid"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
)

type MediaFile struct {
	File         *multipart.FileHeader `form:"file"`
	DocumentType string                `form:"documentType"`
}

func Upload(ctx *gin.Context) {
	var (
		media        model.Media
		uploadFile   MediaFile
		documentType int
		tok          *token.Token
		err          error
		file         *multipart.FileHeader
	)

	tok, err = authentication.GinContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	if tok.UserId == state.ZERO {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	err = ctx.ShouldBind(&uploadFile)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	file = uploadFile.File
	documentType, err = strconv.Atoi(uploadFile.DocumentType)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.Lambda(err),
		})
		return
	}

	media.FileName = file.Filename
	media.Extension = filepath.Ext(file.Filename)
	media.Xid = xid.New().String()

	mType, err := DetectMimeType(file)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.TypeError,
		})
		return
	}

	if !utils.IsValidFile(mType.String()) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.TypeError,
		})
		return
	}

	if utils.IsValidDocument(mType.String()) {
		if documentType == utils.CV || documentType == utils.Letter {
			err = utils.CreateDocumentThumb(media.Xid, media.Extension, file)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
					Message: errx.Lambda(err),
				})
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.TypeError,
		})
		return
	}

	if utils.IsValidVideo(mType.String()) {
		if documentType != utils.VideoPresentation {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.TypeError,
			})
			return
		}

		err = utils.CreateVideoThumb(media.Xid, media.Extension, file)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.Lambda(err),
			})
			return
		}
	}

	if utils.IsValidImage(mType.String()) {
		if documentType != utils.UserProfileImage {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.Lambda(err),
			})
			return
		}

		err = utils.CreateThumb(media.Xid, media.Extension, file)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: errx.Lambda(err),
			})
			return
		}
	}

	err = ctx.SaveUploadedFile(file, utils.FILE_UPLOAD_DIR+media.Xid+media.Extension)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	media.Id, err = database.InsertOne(media)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}

	err = SetUserMediaDetail(documentType, tok.UserId, media.Xid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, media)

}

func DetectMimeType(file *multipart.FileHeader) (mType *mimetype.MIME, err error) {
	readFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	mType, err = mimetype.DetectReader(readFile)
	if err != nil {
		return nil, err
	}
	return mType, nil
}

func SetUserMediaDetail(documentType int, userId int, xId string) (err error) {
	var (
		userMediaDetail model.UserMediaDetail
	)

	userMediaDetail.OwnerId = userId
	userMediaDetail.DocumentType = documentType
	userMediaDetail.DocumentXid = xId

	_, err = database.InsertOne(userMediaDetail)
	if err != nil {
		return err
	}
	return
}

func GetUserMediaDetail(userId, documentType int) (userMediaDetail model.UserMediaDetail, err error) {
	err = database.Get(&userMediaDetail, `SELECT * FROM user_media_detail WHERE owner_id = ?   AND document_type = ?`, userId, documentType)
	if err != nil {
		return userMediaDetail, err
	}
	return userMediaDetail, err
}

func GetMedia(userId, documentType int) (media model.Media, err error) {
	err = database.Get(&media,
		`SELECT media.*
			FROM media
					 JOIN user_media_detail umd ON media.xid = umd.document_xid
					 JOIN user ON user.id = umd.owner_id
			WHERE umd.owner_id = ? AND umd.document_type = ?`, userId, documentType)
	if err != nil {
		return media, err
	}
	return media, err
}

func GetMediaThumb(userId, documentType int) (media model.MediaThumb, err error) {
	err = database.Get(&media,
		`SELECT mt.*
			FROM media_thumb mt
					 JOIN media ON mt.media_xid = media.xid
					 JOIN user_media_detail umd ON umd.document_xid = media.xid
			WHERE umd.owner_id = ?
			  AND umd.document_type = ?`, userId, documentType)
	if err != nil {
		return media, err
	}
	return media, nil
}

func RemoveUserMediaDetail(userMediaDetail model.UserMediaDetail) (err error) {
	err = database.Delete(userMediaDetail)
	if err != nil {
		return err
	}
	return
}

func RemoveMedia(media model.Media) (err error) {
	err = database.Delete(media)
	if err != nil {
		return err
	}
	return
}

func RemoveMediaThumb(mediaThumb model.MediaThumb) (err error) {
	err = database.Delete(mediaThumb)
	if err != nil {
		return err
	}
	return nil
}
