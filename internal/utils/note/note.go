package note

import "errors"

var (
	UnAuthorizedError = "UnAuthorized"
	ParamsError       = errors.New("failed to parse params")
)

var (
	ParseError            = errors.New("cannot parse the corresponding object")
	TypeError             = errors.New("invalid type of file")
	InvalidEmailError     = errors.New("the email you enter is invalid")
	DuplicateUserError    = errors.New("user already exists")
	LinkUserError         = errors.New("failed to link to  user")
	DuplicateAddressError = errors.New("user address already exists")
)

var (
	DatabaseInsertOperationError = errors.New("error while trying to insert data into database")
	DatabaseDeleteOperationError = errors.New("error while trying to delete data from database")
	DatabaseGetOperationError    = errors.New("error while trying to get data from database")
	DatabaseUpdateOperationError = errors.New("error while trying to update data to database")
)

var StatusNeedPasswordError = errors.New("user need at least a password and an email for login")
