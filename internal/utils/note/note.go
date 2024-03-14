package note

const (
	UnAuthorizedError = "UnAuthorized"
)

const (
	ParseError         = "cannot parse the corresponding object"
	InvalidEmailError  = "the email you enter is invalid"
	DuplicateUserError = "user already exists"
)

const (
	DatabaseInsertOperationError = "error while trying to insert data into database"
	DatabaseDeleteOperationError = "error while trying to delete data from database"
	DatabaseGetOperationError    = "error while trying to get data from database"
	DatabaseUpdateOperationError = "error while trying to update data to database"
)
