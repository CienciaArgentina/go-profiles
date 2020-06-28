package profile

import (
	"errors"

	"github.com/CienciaArgentina/go-profiles/config"
)

// NewUserProfileRepository creates a new repository
func NewUserProfileRepository(config.Configuration) (UserProfileRepository, error) {
	// @TODO implement me
	return nil, errors.New("not implemented")
}
