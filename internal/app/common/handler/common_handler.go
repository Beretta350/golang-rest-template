package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Beretta350/golang-rest-template/pkg/errs"
	"github.com/go-playground/validator/v10"
)

func Respond(w http.ResponseWriter, code int, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response(w, code, src)
}

func Error(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError

	switch e := err.(type) {
	case *errs.CustomError:
		code = e.Code.StatusCode()
	case *validator.ValidationErrors, validator.ValidationErrors:
		customErr := errs.ErrValidatingUser.SetDetail(e)
		code = customErr.Code.StatusCode()
		err = customErr
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response(w, code, err)
}

// Respond is response write to ResponseWriter
func response(w http.ResponseWriter, code int, src interface{}) {
	var body []byte
	var err error

	switch s := src.(type) {
	case nil:
		body = []byte{}
	case string:
		body = []byte(s)
	case []byte:
		if !json.Valid(s) {
			Error(w, fmt.Errorf("error: invalid json"))
			return
		}
		body = s
	case error:
		if body, err = json.Marshal(src); err != nil {
			Error(w, fmt.Errorf("error: failed to parse json"))
			return
		}
	default:
		if body, err = json.Marshal(src); err != nil {
			Error(w, err)
			return
		}
	}

	w.WriteHeader(code)
	w.Write(body)
}
