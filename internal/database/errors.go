package database

import (
	"errors"
)

var NotFoundError = errors.New("row not found in database")
