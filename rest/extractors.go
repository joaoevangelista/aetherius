package rest

// ExtractLocation gets the first location on GeoResponse
func ExtractLocation(gr *GeoResponse) LocationResponse {
	results := gr.Results
	if len(results) >= 1 {
		result := results[0]
		return LocationResponse{Location: result.Geometry.Location}
	}
	return LocationResponse{Location: Location{}, Error: NestedError{Status: gr.Status, ErrorMessage: gr.ErrorMessage}}
}

// ExtractAddress gets the FormattedAddress on GeoResponse and wraps it into a AddressResponse
func ExtractAddress(gr *GeoResponse) AddressResponse {
	results := gr.Results
	if len(results) >= 1 {
		result := results[0]
		return AddressResponse{Address: result.FormattedAddress}
	}
	return AddressResponse{Address: "", Error: NestedError{Status: gr.Status, ErrorMessage: gr.ErrorMessage}}
}
