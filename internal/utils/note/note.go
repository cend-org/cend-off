package note

import "errors"

var (
	UnAuthorizedError = errors.New("unauthorized")
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
var AuthError = errors.New(" email or password incorrect  ")

var (
	ThumbError             = errors.New("error while trying to create thumb")
	SavingError            = errors.New("error while trying to save file into server")
	QrError                = errors.New("error while trying to create qr code")
	ComputeError           = errors.New("error while trying to compute salary value")
	MessageError           = errors.New("error while trying to create message")
	MenuItemError          = errors.New("error while trying to create menu item")
	MenuError              = errors.New("error while trying to create menu ")
	VerificationError      = errors.New("error while trying to create verification code ")
	AuthorizationError     = errors.New("error while trying to create authorization  ")
	GenerateMatriculeError = errors.New("error while trying to generate matricule  ")
)
