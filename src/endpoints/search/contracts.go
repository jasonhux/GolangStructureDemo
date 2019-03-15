package search

//Request is used to bind the search params such as search body, headers etc
type Request struct {
	//The calendar year for travel
	//
	// in: query
	Year string `json:"year"`
	//The traveller numbers
	//
	// in: query
	TravellerNumbers string `json:"travellerNum"`
}

//Response is used as the body of search response
type Response struct {
	CountryName   string  `json:"countryName"`
	SearchRequest Request `json:"searchRequest"`
}
