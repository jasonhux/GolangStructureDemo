package home

import (
	"GolangStructureDemo/src/common"
	"net/http"
)

func NewEndpoint() *common.Endpoint {
	return common.NewEndpoint(Handle, nil, Encode, nil)
}

func Handle(r *http.Request, params common.Parameters) (response common.ResponseType, statusCode int, err error, headers map[string]string) {
	return "Golang Structure Demo", http.StatusOK, nil, nil
}

func Encode(w http.ResponseWriter, httpStatus int, response common.ResponseType, headers map[string]string) error {
	common.EncodeTextResponse(w, httpStatus, response)
	return nil
}
