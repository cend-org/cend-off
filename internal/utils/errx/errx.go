package errx

import (
	"duval/internal/utils/note"
	"fmt"
)

func Lambda(err error) string {
	fmt.Println(err)
	return "Something went wrong"
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
