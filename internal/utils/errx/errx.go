package errx

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

var (
	DummyError = "something went wrong please contact support."
)

// User

var (
	UnAuthorizedError = errors.New("unauthorized")
	UnknownUserError  = errors.New("user unknown")
)

//PASSWORD

var (
	EmptyPasswordError  = errors.New("password cannot be empty")
	PasswordLengthError = errors.New("password length is not valid")
	PasswordHashError   = errors.New("cannot protect your password. Something went wrong")
	PasswordInsertError = errors.New("cannot store your password")
)

//LOGIN

var (
	InvalidEmailError       = errors.New("the email you enter is invalid")
	ToRegisterEmailError    = errors.New("user unknown , please sign up instead")
	StatusNeedPasswordError = errors.New("need at least a password and an email for login")
	IncorrectPasswordError  = errors.New("email or password doesn't match")
)

// REGISTRATION

var (
	DuplicateUserError = errors.New("user already exists")
)

// LINK

var (
	UnknownStudentError = errors.New("unknown student")
	ParentError         = errors.New("unknown parent")
	UlError             = errors.New("parent and student are not linked ")
	EmptyTutorError     = errors.New("Continue without tutor ")
)

// DATABASE
var (
	DbInsertError  = errors.New("error while trying to insert data into database")
	DbDeleteError  = errors.New("error while trying to delete data from database")
	DbGetError     = errors.New("error while trying to get data from databaseS please contact support")
	DbUpdateError  = errors.New("update data to database please contact support")
	DuplicateError = errors.New("data already exist in the database")
)

// GENERAL ERROR

var SupportError = errors.New("something went wrong please contact support")

// ACADEMIC

var (
	CoursePreferenceError  = errors.New("no selected course(s)")
	MissingPreferenceError = errors.New("there is no preference in the database")
	LevelError             = errors.New("academic level undefined")
	UnknownLevelError      = errors.New("user academic level undefined")
)

// Language Resource

var (
	LangError  = errors.New("cannot add language resource , resource_ref is missing")
	MLangError = errors.New("resource language is missing")
)

/*

	Function

*/

func IsDuplicate(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
		return true
	}
	return false
}
