package apperr

import "errors"

var ErrNotFound = errors.New("not found")

var ErrUnauthorized = errors.New("unauthorized")

var ErrBadRequest = errors.New("bad request")
