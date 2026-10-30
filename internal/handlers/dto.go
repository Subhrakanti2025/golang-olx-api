package handlers

import (
	"fmt"
	"strings"
	"uuid"
)

type CreateListingRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	City        string `json:"city"`
}

type SendRequestForCreateListing struct {
	Id uuid.UUID `json:"id"`
}

type ValidatorError struct {
	Field   string
	Message string
}

func (err ValidatorError) Error() string {
	return fmt.Sprintf("%s: %s", err.Field, err.Message)
}

func (re CreateListingRequest) Validator() error {
	if strings.TrimSpace(re.Title) == "" {
		return ValidatorError{
			Field:   "Title",
			Message: "Please provide field value",
		}
	} else if strings.TrimSpace(re.Description) == "" {
		return ValidatorError{
			Field:   "Description",
			Message: "Please provide field value",
		}
	} else if (re.Price) == 0 {
		return ValidatorError{
			Field:   "Price",
			Message: "Please provide field value",
		}
	} else if strings.TrimSpace(re.City) == "" {
		return ValidatorError{
			Field:   "City",
			Message: "Please provide field value",
		}
	}

	return nil
}
