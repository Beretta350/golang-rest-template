package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Beretta350/golang-rest-template/internal/pkg/errs"
)

func HttpError(w http.ResponseWriter, err error) {
	if value, ok := err.(*errs.CustomError); ok {
		w.WriteHeader(value.Code.StatusCode())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(err)
}
