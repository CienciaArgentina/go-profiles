package profile

import (
	"errors"

	"github.com/CienciaArgentina/go-profiles/domain"
)

type userProfileService struct {
	repo UserProfileRepository
}

// NewUserProfileService creates a new user profile service
func NewUserProfileService(repo UserProfileRepository) UserProfileService {
	return &userProfileService{repo}
}

func (u *userProfileService) FindUserProfile(id int) (domain.UserProfile, error) {
	return u.repo.Get(id)
}

func (u *userProfileService) GetOrUpdateUserProfile(domain.UserProfile) error {
	// @TODO: implement me
	return errors.New("method not implemented")
}
