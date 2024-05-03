package mediafile

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"io"
	"log"
	"os"
)

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
	err = database.Get(&userMediaDetail,
		`SELECT user_media_detail.* 
			FROM  user_media_detail 
			WHERE user_media_detail.owner_id =?
			  AND user_media_detail.document_type = ?`, userId, documentType)
	if err != nil {
		return userMediaDetail, err
	}
	return userMediaDetail, err
}

func GetMedia(userId, documentType int) (media model.Media, err error) {
	err = database.Get(&media, `SELECT media.*
FROM media
         JOIN user_media_detail ON media.xid = user_media_detail.document_xid
         JOIN user ON user.id = user_media_detail.owner_id
WHERE user_media_detail.owner_id = ? AND user_media_detail.document_type = ?`, userId, documentType)
	if err != nil {
		return media, err
	}
	return media, err
}

func GetMediaThumb(userId, documentType int) (media model.MediaThumb, err error) {
	err = database.Get(&media,
		`SELECT media_thumb.*
				FROM media_thumb
						 JOIN media ON media_thumb.media_xid = media.xid
						 JOIN user_media_detail ON user_media_detail.document_xid = media.xid
				WHERE user_media_detail.owner_id = ?
				  AND user_media_detail.document_type = ?`, userId, documentType)
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

func SaveFile(uploadPath string, file graphql.Upload) error {
	f, err := os.Create(uploadPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = io.Copy(f, file.File)
	if err != nil {
		return err
	}

	err = f.Sync()
	if err != nil {
		return err
	}

	return nil
}
