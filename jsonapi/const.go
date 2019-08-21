package jsonapi

const (
	ERROR_MESSAGE_TOKEN          string = "Invalid authorization token."
	ERROR_MESSAGE_FORBIDDEN      string = "Access forbidden."
	ERROR_MESSAGE_DEFAULT        string = "Internal server error."
	ERROR_MESSAGE_BAD_REQUEST    string = "Bad request."
	ERROR_MESSAGE_DATA_NOT_FOUND string = "Data not found."
	ERROR_MESSAGE_PAGINATION     string = "Failed parsing url pagination."
)

const (
	DEFAULT_LIMIT int = 10
	DEFAULT_PAGE  int = 1
)
