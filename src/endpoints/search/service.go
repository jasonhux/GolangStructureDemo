package search

import (
	"math/rand"
	"net/http"
)

//Service interface is created to allow service injection as long as the used service honours the same function signature
type Service interface {
	Get(params Request) (response Response, status int, e error, headers map[string]string)
}

type service struct {
	//in real world, you may get a country list from storage or upstream service. You can inject the http service to make service test easier
	CountryList []string
}

//NewService is to create a new service
func NewService(countryList []string) Service {
	return &service{countryList}
}

//Get is a function to return recommended country name with provided request params
func (s *service) Get(params Request) (response Response, status int, e error, headers map[string]string) {
	randomIndex := rand.Intn(len(s.CountryList) - 1)
	//returned headers; you may want to return headers based on search results
	headers = map[string]string{"session-id": "xyz"}
	return Response{s.CountryList[randomIndex], params}, http.StatusOK, nil, headers
}
