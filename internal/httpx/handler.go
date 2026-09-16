package httpx

import (
	"errors"
	"net/http"

	"github.com/minwhal/minwhal/internal/apperr"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		status := http.StatusInternalServerError

		switch {
		case errors.Is(err, apperr.ErrBadRequest):
			status = http.StatusBadRequest
		case errors.Is(err, apperr.ErrNotFound):
			status = http.StatusNotFound
		case errors.Is(err, apperr.ErrUnauthorized):
			status = http.StatusUnauthorized
		}
		ResponseWithError(w, status, err)
	}
}
