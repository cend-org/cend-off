package errx

import (
	"duval/internal/utils/note"
	"errors"
	"fmt"
)

func Lambda(err error) error {
	err = errors.New(fmt.Sprintf("Something went wrong : %v", err))
	return err
}

var UnAuthorizedError = note.UnAuthorizedError

var (
	ParseError         = note.ParseError
	InvalidEmailError  = note.InvalidEmailError
	DuplicateUserError = note.DuplicateUserError
)

var (
	DbInsertError = note.DatabaseInsertOperationError
	DbGEtError    = note.DatabaseGetOperationError
	DbDeleteError = note.DatabaseDeleteOperationError
	DbUpdateError = note.DatabaseUpdateOperationError
)
