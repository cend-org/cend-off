package media

import (
	"duval/internal/authentication"
	"duval/internal/configuration"
	"duval/internal/utils"
	"duval/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/joinverse/xid"
	"net/http"
	"path/filepath"
	"time"
)

type Media struct {
	Id          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	FileName    string     `json:"file_name"`
	Extension   string     `json:"extension"`
	Xid         string     `json:"xid"`
	UserId      uint       `json:"user_id"`
	ContentType uint       `json:"content_type"`
}

func Upload(ctx *gin.Context) {
	var (
		media Media
		tok   *authentication.Token
		err   error
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	media.FileName = file.Filename
	media.Extension = filepath.Ext(file.Filename)
	media.Xid = xid.New().String()
	media.UserId = tok.UserId

	err = ctx.SaveUploadedFile(file, utils.FILE_UPLOAD_DIR+media.Xid+media.Extension)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	_, err = database.InsertOne(media)
	if err != nil {
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func ProfileImage(ctx *gin.Context) {
	var (
		err   error
		media Media
		tok   *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	err = database.Get(&media, `SELECT * FROM media WHERE user_id = ? AND content_type = 0`, tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	networkLink := "http://" + configuration.App.Host + ":" + configuration.App.Port + "/api/public/" + media.Xid + media.Extension

	ctx.JSON(http.StatusOK, networkLink)
	return
}
