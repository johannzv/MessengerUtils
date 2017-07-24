package validaters

import (
	"errors"

	"github.com/johannzv/MessengerUtils/models"
)

//ValidateNewUser validates value of new user
func ValidateNewUser(newUser *models.User) error {
	if newUser.PSID == "" {
		return errors.New("missing PSID")
	}
	if newUser.FirstName == "" {
		return errors.New("missing firstName")
	}
	if newUser.LastName == "" {
		return errors.New("missing lastName")
	}
	if newUser.ProfilePic == "" {
		return errors.New("missing profilePic")
	}
	if newUser.Gender == "" {
		return errors.New("missing gender")
	}
	return nil
}
