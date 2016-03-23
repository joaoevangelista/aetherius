package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	APIError        = "API_ERROR"
	JSONEncodeError = "JSON_ENCODE_ERROR"
	EncodeError     = "ENCODE_ERROR"
)

// PrintError send a json error message to writer
func PrintError(w http.ResponseWriter, status string, message string) {
	error := NestedError{Status: status, ErrorMessage: message}
	js, err := json.Marshal(error)
	if err != nil {
		log.Fatal("Error while encoding the error json ")
	}
	fmt.Fprintf(w, string(js))
}

//ThrowAPIErrorIfPresent ...
func ThrowAPIErrorIfPresent(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		PrintError(w, APIError, fmt.Sprintf("Error while decoding %v %v ", err))
	}
}

//ThrowJSONErrorIfPresent ...
func ThrowJSONErrorIfPresent(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		PrintError(w, JSONEncodeError, fmt.Sprintf("Error while encoding json %v", err))
	}
}

// ThrowEncodeErrorIfPresent ...
func ThrowEncodeErrorIfPresent(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		PrintError(w, EncodeError, fmt.Sprintf("Error while encoding number %v", err))
	}
}
