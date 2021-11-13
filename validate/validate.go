package validate

import "net/mail"

const (
	minPasswordLength = 8
	maxPasswordLength = 64

	emptyStringError = "Empty string error"
	tooShortPasswordError = "Too short password error"
	tooLongPasswordError = "Too long password error"
	emailFormatError = "Email format error"
)

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

func Email(email string) error{
		_, err := mail.ParseAddress(email)
		if err != nil{
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