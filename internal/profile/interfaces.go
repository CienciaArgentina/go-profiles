package profile

import (
	"github.com/CienciaArgentina/go-profiles/domain"
	"github.com/gin-gonic/gin"
)

// UserProfileRepository manages CRUD operations
type UserProfileRepository interface {
	Get(int) (domain.UserProfile, error)
	Create(domain.UserProfile) error
	Update(domain.UserProfile) error
	Delete(int) error
}

// UserProfileService defines the user profile related API
type UserProfileService interface {
	FindUserProfile(int) (domain.UserProfile, error)
	CreateUserProfile(domain.UserProfile) error
	UpdateUserProfile(int, domain.UserProfile) error
	DeleteUserProfile(int) error
}

// UserProfileController defines the user profile controller
type UserProfileController interface {
	Get(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}
