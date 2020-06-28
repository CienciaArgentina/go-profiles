package profile

import (
	"errors"

	"github.com/CienciaArgentina/go-profiles/config"
)

// NewUserProfileService creates a new user profile service
func NewUserProfileService(config.Configuration) (UserProfileService, error) {
	// @TODO implement me
	return nil, errors.New("not implemented")
}
