package errorhandling

import "errors"

var (
	ErrBadRequest     = errors.New("invalid request data or parameters")
	ErrUnauthorized   = errors.New("authorization required")
	ErrForbidden      = errors.New("request denied")
	ErrNotFound       = errors.New("resource not found")
	ErrInternalServer = errors.New("server encountered an error")
)
