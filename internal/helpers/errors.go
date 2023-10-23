package helpers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var ErrInvalidPathParam = errors.New("invalid path param")

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch err {
	case ErrInvalidPathParam:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
