package profile

import (
	"github.com/CienciaArgentina/go-profiles/domain"
	"github.com/gin-gonic/gin"
)

// UserProfileRepository manages CRUD operations
type UserProfileRepository interface {
	Get(int)
	Create(domain.UserProfile)
	Update(domain.UserProfile)
	Delete(int)
}

// UserProfileService defines the user profile related API
type UserProfileService interface {
	FindUserProfile(int)
	GetOrUpdateUserProfile(domain.UserProfile)
}

// UserProfileController defines the user profile controller
type UserProfileController interface {
	Get(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}
