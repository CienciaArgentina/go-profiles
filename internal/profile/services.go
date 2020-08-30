package profile

import (
	"github.com/CienciaArgentina/go-profiles/domain"
	"github.com/CienciaArgentina/go-profiles/internal/errors"
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

func (u *userProfileService) CreateUserProfile(userProfile domain.UserProfile) error {
	if _, err := u.repo.Get(userProfile.UserID); err != nil {
		if err == errors.ErrUserProfileNotFound {
			return u.repo.Create(userProfile)
		}
		return err
	}

	return errors.ErrUserProfileAlreadyExists
}

func (u *userProfileService) UpdateUserProfile(id int, userProfile domain.UserProfile) error {
	if _, err := u.repo.Get(id); err != nil {
		if err != errors.ErrUserProfileNotFound {
			return err
		}
	}
	userProfile.UserID = id
	return u.repo.Update(userProfile)
}

func (u *userProfileService) DeleteUserProfile(id int) error {
	return u.repo.Delete(id)
}
