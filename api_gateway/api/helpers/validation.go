package helpers

import (
	"errors"
)

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be blank")
	}
	if len(password) < 4 || len(password) > 30 {
		return errors.New("password length should be 4 to 30 characters")
	}
	return nil
}

func ValidatePhoneNumber(phoneNumber string) error {
	if phoneNumber == "" {
		return errors.New("phoneNumber cannot be blank")
	}
	if len(phoneNumber) != 13 {
		return errors.New("phoneNumber length should be 13. Example +998.........")
	}
	return nil
}
