package authentication

import (
	"context"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/configuration"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/joinverse/xid"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"strconv"
	"time"
)

func GenerateQrCode(ctx context.Context) (*string, error) {
	var (
		tok            *token.Token
		err            error
		networkLink    string
		qrImageLink    string
		QRCodeRegistry model.QRCodeRegistry
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &qrImageLink, errx.UnAuthorizedError
	}

	networkLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/api/login/with-qr/:" + strconv.Itoa(int(tok.UserId))

	qrc, err := qrcode.New(networkLink)
	if err != nil {
		return &qrImageLink, errx.QrError
	}

	QRCodeRegistry.UserId = tok.UserId
	QRCodeRegistry.Xid = xid.New().String()
	QRCodeRegistry.IsUsed = false

	qrImageLink = "http://" + configuration.App.Host + ":" + configuration.App.Port + "/api/public/qr/" + QRCodeRegistry.Xid + ".jpg"

	w, err := standard.New(utils.FILE_UPLOAD_DIR + utils.QR_CODE_UPLOAD_DIR + QRCodeRegistry.Xid + ".jpg")
	if err != nil {
		return &qrImageLink, errx.SavingError

	}
	// save file
	err = qrc.Save(w)
	if err != nil {
		return &qrImageLink, errx.SavingError
	}
	_, err = database.InsertOne(QRCodeRegistry)
	if err != nil {
		return &qrImageLink, errx.DbInsertError

	}

	return &qrImageLink, nil
}

func LoginWithQr(ctx context.Context, xId string) (*string, error) {
	var (
		tok         string
		err         error
		qrCode      model.QRCodeRegistry
		currentUser model.User
	)

	qrCode, err = GetQRCodeRegistry(xId)
	if err != nil {
		return &tok, errx.DbGetError

	}

	err = UpdateQRCodeRegistryFlag(qrCode)
	if err != nil {
		return &tok, nil
	}

	currentUser, err = GetUserWithId(qrCode.UserId)
	if err != nil {
		return &tok, errx.UnAuthorizedError
	}

	tok, err = LoginWithEmail(currentUser.Email)
	if err != nil {
		return &tok, nil
	}
	return &tok, nil
}

/*
	UTILS
*/

func GetQRCodeRegistry(xId string) (QRCodeRegistry model.QRCodeRegistry, err error) {
	err = database.Get(&QRCodeRegistry, `SELECT * FROM qr_code_registry WHERE qr_code_registry.xid = ?`, xId)
	if err != nil {
		return QRCodeRegistry, err
	}
	return QRCodeRegistry, nil
}

func UpdateQRCodeRegistryFlag(QRCodeRegistry model.QRCodeRegistry) (err error) {
	QRCodeRegistry.IsUsed = true
	err = database.Update(QRCodeRegistry)
	if err != nil {
		return err
	}
	return
}

func GetUserWithId(id int) (user model.User, err error) {
	time.Sleep(100)
	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, id)
	if err != nil {
		return user, err
	}

	return user, err
}

func LoginWithEmail(email string) (access string, err error) {
	var usr model.User

	err = database.Get(&usr, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return
	}

	access, err = NewAccessToken(usr)
	if err != nil {
		return access, err
	}

	return access, err
}
