package search

import (
	"GolangStructureDemo/src/common"
	"net/http"
)

//NewEndpoint is to generate a search endpoint
func NewEndpoint(s Service) *common.Endpoint {
	return common.NewEndpoint(Handle(s), Bind, Encode, nil)
}

//Bind is to bind a search request with given http request
func Bind(r *http.Request) (params common.Parameters, err error) {

	queryMap := r.URL.Query()
	searchParams := Request{
		TravellerNumbers: getQueryParam(queryMap, "travellerNum"),
		Year:             getQueryParam(queryMap, "year"),
	}
	//Consider to add a validation on the binded searchParams before return
	return searchParams, nil

}

//Encode is to encode json response for search endpoint; here AddHeadersToResponse is included for demostration as well
func Encode(w http.ResponseWriter, httpStatus int, response common.ResponseType, headers map[string]string) error {
	common.AddHeadersToResponse(w, headers)
	common.EncodeJsonResponse(w, httpStatus, response)
	return nil
}

//Handle is to create response based on the search request
func Handle(s Service) func(r *http.Request, params common.Parameters) (response common.ResponseType, statusCode int, err error, headers map[string]string) {
	return func(r *http.Request, params common.Parameters) (response common.ResponseType, statusCode int, err error, headers map[string]string) {
		return s.Get(params.(Request))
	}
}

func getQueryParam(queryMap map[string][]string, paramKey string) (value string) {
	values, ok := queryMap[paramKey]
	if !ok || len(values[0]) < 1 {
		return ""
	}
	return values[0]
}
