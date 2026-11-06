package httpx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type code string

type codeStruct struct {
	Value code
}

var (
	invalidID           = codeStruct{Value: "invalid_id"}
	listingNotFound     = codeStruct{Value: "listing_not_found"}
	badRequest          = codeStruct{Value: "invalid_resquest"}
	validationError     = codeStruct{Value: "validation_error"}
	internalServerError = codeStruct{Value: "internal_server_error"}
	statusConflict      = codeStruct{Value: "field_conflict"}
)

func InvalidID() codeStruct {
	return invalidID
}
func ListingNotFound() codeStruct {
	return listingNotFound
}
func BadRequest() codeStruct {
	return badRequest
}
func ValidationError() codeStruct {
	return validationError

}
func InternalServerError() codeStruct {
	return internalServerError

}
func StatusConflict() codeStruct {
	return statusConflict
}

type errorEnvelope struct {
	Error errorPayload `json:"error"`
}
type errorPayload struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Field   string `json:"field,omitempty"`
}

func Error(w http.ResponseWriter, statusCode int, errorMessage string, code codeStruct) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	error := errorEnvelope{
		Error: errorPayload{
			Message: errorMessage,
			Code:    string(code.Value),
		},
	}
	json.NewEncoder(w).Encode(error)
}

func ValidatorError(w http.ResponseWriter, statusCode int, errorMessage string, code codeStruct, field string) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("statusCode", statusCode)
	w.WriteHeader(statusCode)
	error := errorEnvelope{
		Error: errorPayload{
			Message: errorMessage,
			Code:    string(code.Value),
			Field:   field,
		},
	}
	json.NewEncoder(w).Encode(error)
}
