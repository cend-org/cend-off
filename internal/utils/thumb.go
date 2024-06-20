package utils

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/disintegration/imaging"
	"github.com/gen2brain/svg"
	"github.com/joinverse/xid"
	mod "github.com/unidoc/unipdf/v3/model"
	"github.com/unidoc/unipdf/v3/render"
	"golang.org/x/image/webp"
	"image"
	"image/color"
	"mime/multipart"
)

/*
CREATE THUMBNAIL FOR UPLOADED IMAGE
*/
func CreateThumb(mediaXid string, mType string, file *multipart.FileHeader) (err error) {
	var (
		mediaThumb model.MediaThumb
		thumbnail  image.Image
		img        image.Image
	)

	openedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer openedFile.Close()

	if IsValidWebp(mType) {
		img, err = webp.Decode(openedFile)

	} else if IsValidSvg(mType) {
		img, err = svg.Decode(openedFile)
	} else {
		img, err = imaging.Decode(openedFile)
	}
	if err != nil {
		return err
	}

	mediaThumb.Extension = ".jpg"
	mediaThumb.MediaXid = mediaXid
	mediaThumb.Xid = "T_" + xid.New().String()

	thumbnail = imaging.Thumbnail(img, 200, 200, imaging.CatmullRom)

	dst := imaging.New(200, 200, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, thumbnail, image.Pt(0, 0))
	err = imaging.Save(dst, FILE_UPLOAD_DIR+THUMB_FILE_UPLOAD_DIR+mediaThumb.Xid+".jpg")
	if err != nil {
		return
	}

	_, err = database.InsertOne(mediaThumb)
	if err != nil {
		return err
	}

	return
}

/*
CREATE THUMBNAIL FOR UPLOADED COVER LETTER
*/
func CreateDocumentThumb(mediaXid string, extension string, file *multipart.FileHeader) (err error) {
	var (
		mediaThumb model.MediaThumb
		thumbnail  image.Image
	)

	openedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer openedFile.Close()

	pdfReader, err := mod.NewPdfReader(openedFile)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil || numPages < 1 {
		return err
	}

	firstPage, err := pdfReader.GetPage(1)
	if err != nil {
		return err
	}
	device := render.NewImageDevice()

	img, err := device.Render(firstPage)
	if err != nil {
		return err
	}

	mediaThumb.Extension = ".jpg"
	mediaThumb.MediaXid = mediaXid
	mediaThumb.Xid = "T_" + xid.New().String()

	thumbnail = imaging.Thumbnail(img, 800, 1100, imaging.CatmullRom)

	dst := imaging.New(800, 1100, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, thumbnail, image.Pt(0, 0))
	err = imaging.Save(dst, FILE_UPLOAD_DIR+THUMB_FILE_UPLOAD_DIR+mediaThumb.Xid+".jpg")
	if err != nil {
		return err
	}

	_, err = database.InsertOne(mediaThumb)
	if err != nil {
		return err
	}
	return
}

/*
CREATE THUMBNAIL FOR UPLOADED Video
*/

func CreateVideoThumb(mediaXid string, file *multipart.FileHeader) (err error) {
	var (
		mediaThumb model.MediaThumb
	)

	mediaThumb.Extension = ".jpg"
	mediaThumb.MediaXid = mediaXid
	mediaThumb.Xid = "T_" + xid.New().String()

	_, err = database.InsertOne(mediaThumb)
	if err != nil {
		return err
	}
	return
}
