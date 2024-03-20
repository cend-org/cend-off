package thumb

import (
	"duval/internal/utils"
	"duval/pkg/database"
	"image"
	"image/color"
	"mime/multipart"
	"time"

	"github.com/disintegration/imaging"
	"github.com/joinverse/xid"
)

type MediaThumb struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Extension string     `json:"extension"`
	MediaXid  string     `json:"media_xid"`
	Xid       string     `json:"xid"`
}

/*
CREATE THUMBNAIL FOR UPLOADED IMAGE
*/
func CreateThumb(mediaXid string, extension string, file multipart.File) (err error) {
	var (
		mediaThumb MediaThumb
		thumbnail  image.Image
	)

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

	mediaThumb.Extension = extension
	mediaThumb.MediaXid = mediaXid
	mediaThumb.Xid = "T_" + xid.New().String()

	_, err = database.InsertOne(mediaThumb)
	if err != nil {
		return err
	}

	return
}
