package authentication

import (
	"duval/internal/configuration"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/joinverse/xid"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"net/http"
	"strconv"
	"time"
)

type QrCodeRegistry struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	UserId    uint       `json:"user_id"`
	Xid       string     `json:"xid"`
	IsUsed    bool       `json:"is_used"`
}

func GenerateQrCode(ctx *gin.Context) {
	var (
		tok            *Token
		err            error
		networkLink    string
		qrImageLink    string
		qrCodeRegistry QrCodeRegistry
	)
	time.Sleep(100)
	tok, err = GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/api/login/with-qr/:" + strconv.Itoa(int(tok.UserId))

	qrc, err := qrcode.New(networkLink)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	qrCodeRegistry.UserId = tok.UserId
	qrCodeRegistry.Xid = xid.New().String()
	qrCodeRegistry.IsUsed = false

	qrImageLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/api/public/qr/" + qrCodeRegistry.Xid + ".jpg"

	w, err := standard.New(utils.FILE_UPLOAD_DIR + utils.QR_CODE_UPLOAD_DIR + qrCodeRegistry.Xid + ".jpg")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}
	// save file
	err = qrc.Save(w)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}
	_, err = database.InsertOne(qrCodeRegistry)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, qrImageLink)
}
