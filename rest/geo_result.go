package rest

//GeoResponse wrapper for reponse
type GeoResponse struct {
	Results      []Result `json:"results"`
	Status       string   `json:"status"`
	ErrorMessage string   `json:"error_message"`
}

// Result a wrapper for adresses
type Result struct {
	AddressComponents []Address `json:"address_components"`
	FormattedAddress  string    `json:"formatted_address"`
	Geometry          Geometry  `json:"geometry"`
	PlaceID           string    `json:"place_id"`
	Types             []string  `json:"types"`
}

// Address contain human readable addresses
type Address struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

//Geometry Maps locations
type Geometry struct {
	Bounds       map[string]Location `json:"bounds"`
	Location     Location            `json:"location"`
	LocationType string              `json:"location_type"`
	Viewport     map[string]Location `json:"viewport"`
}

// Location Latitude and Longitude
type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}
