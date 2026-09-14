package database

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("not found")

func WrapNotFound(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}
	return err
}
