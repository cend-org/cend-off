package note

const (
	UnAuthorizedError = "UnAuthorized"
	ParamsError       = "failed to parse params"
)

const (
	ParseError            = "cannot parse the corresponding object"
	TypeError             = "invalid type of file"
	InvalidEmailError     = "the email you enter is invalid"
	DuplicateUserError    = "user already exists"
	LinkUserError         = "failed to link to  user"
	DuplicateAddressError = "user address already exists"
)

const (
	DatabaseInsertOperationError = "error while trying to insert data into database"
	DatabaseDeleteOperationError = "error while trying to delete data from database"
	DatabaseGetOperationError    = "error while trying to get data from database"
	DatabaseUpdateOperationError = "error while trying to update data to database"
)

const StatusNeedPasswordError = "user need at least a password and an email for login"
