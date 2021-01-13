package validation

import "errors"

var (
	// ErrInvalidName invalid name error
	ErrInvalidName = errors.New("invalid name")
	// ErrInvalidEmail invalid email error
	ErrInvalidEmail = errors.New("invalid email")
	// ErrInvalidPassword invalid password error
	ErrInvalidPassword = errors.New("invalid password")
)

// UserName validates if name provided is valid
func UserName(name string) error {
	if len(name) < 3 {
		return ErrInvalidName
	}
	return nil
}

// UserEmail validates if email provided is valid
func UserEmail(email string) error {
	if len(email) < 3 {
		return ErrInvalidEmail
	}
	return nil
}

// UserPassword validates if password provided is valid
func UserPassword(password string) error {
	if len(password) < 3 {
		return ErrInvalidPassword
	}
	return nil
}
