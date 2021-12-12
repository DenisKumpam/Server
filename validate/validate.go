package validate

import (
	"log"
	"net/mail"
)

const (
	minPasswordLength = 8
	maxPasswordLength = 64

	emptyStringError      = "Empty string error"
	tooShortPasswordError = "Too short password error"
	tooLongPasswordError  = "Too long password error"
	emailFormatError      = "Email format error"
	NameField             = "name"
)

type Form struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type validationErr struct {
	Fail string `json:"key"`
	Err  string `json:"Err"`
}

func (ve validationErr) Error() string {
	return ve.Err
}

func (fd *Form) ValidateForm() []error {
	errors := make([]error, 0)

	if Empty(fd.Name) {
		errors = append(errors, validationErr{NameField, emptyStringError})
	}

	err := Email(fd.Email)
	if err != nil {
		errors = append(errors, err)
	}

	err = Password(fd.Password)
	if err != nil {
		errors = append(errors, err)
	}

	return errors
}

func Empty(str string) bool {
	if str == "" {
		return true
	}
	return false
}

func Email(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Printf("%s\n", err.Error())
		return validationErr{"Email", emailFormatError}
	}
	return nil
}

func Password(password string) error {
	if password == "" {
		return &validationErr{"Pass", emptyStringError}
	}

	if len(password) < minPasswordLength {
		return &validationErr{"ShortPass", tooShortPasswordError}
	}

	if len(password) > maxPasswordLength {
		return &validationErr{"LongPass", tooLongPasswordError}
	}

	return nil
}
