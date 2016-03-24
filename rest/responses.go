package rest

//structs here are pointers so it will not be serialized if not set

// AddressResponse wraps a json reponse for single address
type AddressResponse struct {
	Address *string      `json:"address,omitempty"`
	Error   *NestedError `json:"errors,omitempty"`
}

// LocationResponse wraps a json response for a single location
type LocationResponse struct {
	Location *Location    `json:"location,omitempty"`
	Error    *NestedError `json:"error,omitempty"`
}

// NestedError wraps a error struc to inform any errors occured when communiting with API
type NestedError struct {
	Status       string `json:"status,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type Health struct {
	Status string `json:"status"`
}
