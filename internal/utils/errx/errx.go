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

var UnAuthorizedError = errors.New(note.UnAuthorizedError)

var (
	ParseError         = errors.New(note.ParseError)
	InvalidEmailError  = errors.New(note.InvalidEmailError)
	DuplicateUserError = errors.New(note.DuplicateUserError)
)

var (
	DbInsertError = errors.New(note.DatabaseInsertOperationError)
	DbDeleteError = errors.New(note.DatabaseDeleteOperationError)
)
