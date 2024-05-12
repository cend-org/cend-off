package errx

import (
	"errors"
	"fmt"
	"github.com/cend-org/duval/internal/utils/note"
)

func Lambda(err error) error {
	err = errors.New(fmt.Sprintf("Something went wrong : %v", err))
	return err
}

var UnAuthorizedError = note.UnAuthorizedError
var ParamsError = note.ParamsError

var (
	ParseError            = note.ParseError
	TypeError             = note.TypeError
	InvalidEmailError     = note.InvalidEmailError
	DuplicateUserError    = note.DuplicateUserError
	LinkUserError         = note.LinkUserError
	DuplicateAddressError = note.DuplicateAddressError
)

var (
	DbInsertError = note.DatabaseInsertOperationError
	DbGetError    = note.DatabaseGetOperationError
	DbDeleteError = note.DatabaseDeleteOperationError
	DbUpdateError = note.DatabaseUpdateOperationError
)

var (
	NeedPasswordError = note.StatusNeedPasswordError
)

var (
	ThumbError             = note.ThumbError
	SavingError            = note.SavingError
	QrError                = note.QrError
	ComputeError           = note.ComputeError
	MessageError           = note.MessageError
	MenuItemError          = note.MenuItemError
	MenuError              = note.MenuError
	VerificationError      = note.VerificationError
	AuthorizationError     = note.AuthorizationError
	GenerateMatriculeError = note.GenerateMatriculeError
)
