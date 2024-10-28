package common

import "net/http"

type ErrorCode struct {
	Name       string
	StatusCode int
}

/*
	General error codes
*/

var ErrorCodeInternalProcess = ErrorCode{
	Name:       "INTERNAL_PROCESS",
	StatusCode: http.StatusInternalServerError,
}

/*
	Authentication and Authorization error codes
*/

var ErrorCodeAuthPermissionDenied = ErrorCode{
	Name:       "AUTH_PERMISSION_DENIED",
	StatusCode: http.StatusForbidden,
}

var ErrorCodeAuthNotAuthenticated = ErrorCode{
	Name:       "AUTH_NOT_AUTHENTICATED",
	StatusCode: http.StatusUnauthorized,
}

/*
	Resource-related error codes
*/

// Data not found
var ErrorCodeResourceNotFound = ErrorCode{
	Name:       "RESOURCE_NOT_FOUND",
	StatusCode: http.StatusNotFound,
}

// Data duplicated
var ErrorCodeResourceConflict = ErrorCode{
	Name:       "RESOURCE_CONFLICT",
	StatusCode: http.StatusConflict,
}

/*
	Parameter-related error codes
*/

// Invalid parameter
var ErrorCodeParameterInvalid = ErrorCode{
	Name:       "PARAMETER_INVALID",
	StatusCode: http.StatusBadRequest,
}

/*
	Remote server-related error codes
*/

var ErrorCodeRemoteProcess = ErrorCode{
	Name:       "REMOTE_PROCESS_ERROR",
	StatusCode: http.StatusBadGateway,
}
