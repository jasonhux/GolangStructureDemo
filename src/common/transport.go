package common

import (
	"net/http"
)

type Parameters interface{}
type ResponseType interface{}
type Handler func(r *http.Request, params Parameters) (response ResponseType, statusCode int, err error, headers map[string]string)
type Binder func(r *http.Request) (params Parameters, err error)
type ResponseEncoder func(w http.ResponseWriter, httpStatus int, response ResponseType, headers map[string]string) error
type ErrorLogger func(r *http.Request, httpStatus int, e error)

type Endpoint struct {
	handler Handler
	binder  Binder
	encoder ResponseEncoder
	logger  ErrorLogger
}

func NewEndpoint(handler Handler, binder Binder, encoder ResponseEncoder, logger ErrorLogger) *Endpoint {
	return &Endpoint{handler, binder, encoder, logger}
}

func GetHandler(api *Endpoint) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var params Parameters
		var e error
		var resp ResponseType
		var statusCode int
		var headers map[string]string

		if api.binder != nil {
			params, e = api.binder(r)
			if e != nil {
				statusCode = http.StatusBadRequest
			}
		}
		if e == nil {
			resp, statusCode, e, headers = api.handler(r, params)
		}

		if e != nil {
			api.logger(r, statusCode, e)
			newErrorResponse := ErrorResponse{e.Error()}
			api.encoder(w, statusCode, newErrorResponse, headers)
		} else {
			api.encoder(w, statusCode, resp, headers)
		}
	})
}
