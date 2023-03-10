package errors

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : article
 */

const (
	UNKNOWN_ERROR             ErrorCode = 0
	DATA_INVALID              ErrorCode = 4_001_00_00001
	FAILED_RETRIEVE_DATA      ErrorCode = 4_001_00_00002
	STATUS_PAGE_NOT_FOUND     ErrorCode = 4_001_00_00003
	INVALID_HEADER_X_PLATFORM ErrorCode = 4_001_00_00004
	INVALID_HEADER_X_API_KEY  ErrorCode = 4_001_00_00005
	UNAUTHORIZED              ErrorCode = 4_001_00_00006
	FAILED_CREATE_DATA        ErrorCode = 4_001_00_00007
)

var errorCodes = map[ErrorCode]*CommonError{
	UNKNOWN_ERROR: {
		ClientMessage: "Unknown error.",
		SystemMessage: "Unknown error.",
		ErrorCode:     UNKNOWN_ERROR,
	},
	DATA_INVALID: {
		ClientMessage: "Invalid Data Request",
		SystemMessage: "Some of query params has invalid value.",
		ErrorCode:     DATA_INVALID,
	},
	FAILED_RETRIEVE_DATA: {
		ClientMessage: "Failed to retrieve Data.",
		SystemMessage: "Something wrong happened while retrieve Data.",
		ErrorCode:     FAILED_RETRIEVE_DATA,
	},
	STATUS_PAGE_NOT_FOUND: {
		ClientMessage: "Invalid Status Page.",
		SystemMessage: "Status Page Email Address not found.",
		ErrorCode:     STATUS_PAGE_NOT_FOUND,
	},
	INVALID_HEADER_X_PLATFORM: {
		ClientMessage: "Invalid platform.",
		SystemMessage: "Invalid value of header X-Platform.",
		ErrorCode:     INVALID_HEADER_X_PLATFORM,
	},
	INVALID_HEADER_X_API_KEY: {
		ClientMessage: "Invalid Api Key.",
		SystemMessage: "Invalid value of header X-Api-Key.",
		ErrorCode:     INVALID_HEADER_X_API_KEY,
	},
	UNAUTHORIZED: {
		ClientMessage: "Unauthorized",
		SystemMessage: "Unauthorized",
		ErrorCode:     UNAUTHORIZED,
	},
	FAILED_CREATE_DATA: {
		ClientMessage: "Failed to create data.",
		SystemMessage: "Something wrong happened while create data.",
		ErrorCode:     FAILED_CREATE_DATA,
	},
}
