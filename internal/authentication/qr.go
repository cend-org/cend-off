package authentication

import (
	"context"
	"duval/internal/configuration"
	"duval/internal/graph/model"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/pkg/database"
	"github.com/joinverse/xid"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"strconv"
)

func GenerateQrCode(ctx *context.Context) (*string, error) {
	var (
		tok            *Token
		err            error
		networkLink    string
		qrImageLink    string
		qrCodeRegistry model.QrCodeRegistry
	)

	tok, err = GetTokenDataFromContext(*ctx)
	if err != nil {
		return &qrImageLink, errx.UnAuthorizedError
	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/api/login/with-qr/:" + strconv.Itoa(int(tok.UserId))

	qrc, err := qrcode.New(networkLink)
	if err != nil {
		return &qrImageLink, errx.Lambda(err)
	}

	qrCodeRegistry.UserId = tok.UserId
	qrCodeRegistry.Xid = xid.New().String()
	qrCodeRegistry.IsUsed = false

	qrImageLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/api/public/qr/" + qrCodeRegistry.Xid + ".jpg"

	w, err := standard.New(utils.FILE_UPLOAD_DIR + utils.QR_CODE_UPLOAD_DIR + qrCodeRegistry.Xid + ".jpg")
	if err != nil {
		return &qrImageLink, errx.Lambda(err)

	}
	// save file
	err = qrc.Save(w)
	if err != nil {
		return &qrImageLink, errx.Lambda(err)
	}
	_, err = database.InsertOne(qrCodeRegistry)
	if err != nil {
		return &qrImageLink, errx.DbInsertError

	}

	return &qrImageLink, nil
}

func LoginWithQr(ctx *context.Context, xId string) (*string, error) {
	var (
		tok            string
		err            error
		qrCodeRegistry model.QrCodeRegistry
	)

	qrCodeRegistry, err = GetQrCodeRegistry(xId)
	if err != nil {
		return &tok, errx.Lambda(err)

	}

	err = UpdateQrCodeRegistryFlag(qrCodeRegistry)
	if err != nil {
		return &tok, nil
	}

	tok, err = GetTokenString(qrCodeRegistry.UserId)
	if err != nil {
		return &tok, errx.Lambda(err)

	}

	return &tok, nil
}

/*
	UTILS
*/

func GetQrCodeRegistry(xId string) (qrCodeRegistry model.QrCodeRegistry, err error) {
	err = database.Get(&qrCodeRegistry, `SELECT * FROM qr_code_registry WHERE qr_code_registry.xid = ?`, xId)
	if err != nil {
		return qrCodeRegistry, err
	}
	return qrCodeRegistry, nil
}

func UpdateQrCodeRegistryFlag(qrCodeRegistry model.QrCodeRegistry) (err error) {
	qrCodeRegistry.IsUsed = true
	err = database.Update(qrCodeRegistry)
	if err != nil {
		return err
	}
	return
}
