package profile

import (
	"duval/internal/utils"
	"duval/pkg/database"
	"image"
	"image/color"
	"mime/multipart"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gabriel-vasile/mimetype"
	"github.com/joinverse/xid"
)

type Thumb struct {
	Id          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	FileName    string     `json:"file_name"`
	Extension   string     `json:"extension"`
	MediaXid    string     `json:"media_xid"`
	ContentType string     `json:"content_type"`
}

/*
CREATE THUMBNAIL FOR UPLOADED IMAGE
*/
func CreateThumb(mediaXid string, extension string, file multipart.File) (err error) {
	var (
		thumb     Thumb
		thumbnail image.Image
	)

	filePath := utils.FILE_UPLOAD_DIR + mediaXid + extension
	mtype, err := mimetype.DetectFile(filePath)
	if err != nil {
		return err
	}

	img, err := imaging.Decode(file)
	if err != nil {
		return err
	}

	thumbnail = imaging.Thumbnail(img, 200, 200, imaging.CatmullRom)

	dst := imaging.New(200, 200, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, thumbnail, image.Pt(0, 0))
	err = imaging.Save(dst, utils.FILE_UPLOAD_DIR+utils.THUMB_FILE_UPLOAD_DIR+mediaXid+extension)
	if err != nil {
		return
	}

	thumb.ContentType = mtype.String()
	thumb.Extension = extension
	thumb.MediaXid = mediaXid
	thumb.FileName = xid.New().String() + extension

	_, err = database.InsertOne(thumb)
	if err != nil {
		return err
	}

	return
}
