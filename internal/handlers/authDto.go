package handlers

import "fmt"

type UserStruct struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInsertResposne struct {
	Id string `json:"id"`
}

type AuthValidatorErr struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (authErr AuthValidatorErr) Error() string {
	return fmt.Sprintf("%s: %s", authErr.Field, authErr.Message)
}

func (usr UserStruct) AuthValidate() error {
	if usr.Name == "" {
		return &AuthValidatorErr{
			Field:   "Name",
			Message: "Please provide name",
		}
	}
	// validate email
	if usr.Email == "" {
		return &AuthValidatorErr{
			Field:   "Email",
			Message: "Please provide email",
		}
	}

	if len(usr.Password) < 8 {
		return &AuthValidatorErr{
			Field:   "Email",
			Message: "Please provide 8 character password",
		}
	}

	return nil
}
