package api

import (
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrInternalServer   = &ErrorResponse{StatusCode: http.StatusInternalServerError, ErrorText: "internal server error"}
	ErrUnauthorized     = &ErrorResponse{StatusCode: http.StatusUnauthorized, ErrorText: "unauthorized"}
	ErrDecodeJsonFailed = &ErrorResponse{StatusCode: http.StatusBadRequest, ErrorText: "failed to decode JSON"}
)

var _ render.Renderer = &ErrorResponse{}
var _ error = &ErrorResponse{}

type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	ErrorText  string `json:"error"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func (e *ErrorResponse) Error() string {
	if e.Err == nil {
		return e.ErrorText
	}

	return e.Err.Error()
}

func (e *ErrorResponse) WithError(err error) *ErrorResponse {
	return &ErrorResponse{Err: err, StatusCode: e.StatusCode, ErrorText: e.ErrorText}
}

func ErrExecute(err error) *ErrorResponse {
	return &ErrorResponse{Err: err, StatusCode: http.StatusUnprocessableEntity, ErrorText: "failed to process request"}
}
