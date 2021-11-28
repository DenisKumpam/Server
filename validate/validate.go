package validate

import "net/mail"

const (
	minPasswordLength = 8
	maxPasswordLength = 64

	emptyStringError      = "Empty string error"
	tooShortPasswordError = "Too short password error"
	tooLongPasswordError  = "Too long password error"
	emailFormatError      = "Email format error"
)

type FormD map[string]string

type validationErr struct {
	Key string `json:"key"`
	Err string `json:"Err"`
}

type ValidationErrs struct {
	Errors []validationErr
}

type regError struct {
	err string
}

func (r regError) Error() string {
	return r.err
}

func Empty(str string) error {
	if str == "" {
		return regError{emptyStringError}
	}
	return nil
}

func Email(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return regError{emailFormatError}
	}
	return nil
}
func Password(password string) error {
	if password == "" {
		return &regError{emptyStringError}
	}

	if len(password) < minPasswordLength {
		return &regError{tooShortPasswordError}
	}

	if len(password) > maxPasswordLength {
		return &regError{tooLongPasswordError}
	}

	return nil
}
func ValidateFormD(r *FormD) ValidationErrs {
	errors := ValidationErrs{}

	for k, val := range *r {
		err := validation(k, val)
		if err != nil {
			errors.Errors = append(errors.Errors, validationErr{k, err.Error()})
		} else {
			errors.Errors = append(errors.Errors, validationErr{k, ""})
		}
	}
	return errors
}

func validation(key, str string) error {
	err := Empty(str)
	if err != nil {
		return err
	}

	if key == "email" {
		err = Email(str)
	} else {
		if key == "password" {
			err = Password(str)
		}
	}
	return err
}
