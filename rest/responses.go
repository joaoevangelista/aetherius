package rest

// AddressResponse wrapps a json reponse for single address
type AddressResponse struct {
	Address string      `json:"address"`
	Error   NestedError `json:"errors"`
}

// LocationResponse wrapps a json response for a single location
type LocationResponse struct {
	Location Location    `json:"location"`
	Error    NestedError `json:"error"`
}

// NestedError wrapps a error struc to inform any errors occured when communiting with API
type NestedError struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}
