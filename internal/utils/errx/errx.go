package errx

import (
	"duval/internal/utils/note"
	"errors"
)

func Lambda(err error) error {
	return err
}

var (
	ParseError         = errors.New(note.ParseError)
	InvalidEmailError  = errors.New(note.InvalidEmailError)
	DuplicateUserError = errors.New(note.DuplicateUserError)
)

var (
	DbInsertError = errors.New(note.DatabaseInsertOperationError)
	DbDeleteError = errors.New(note.DatabaseDeleteOperationError)
)
