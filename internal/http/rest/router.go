package rest

import (
	"github.com/CienciaArgentina/go-backend-commons/pkg/infrastructure/database"
	"github.com/CienciaArgentina/go-profiles/config"
	"github.com/CienciaArgentina/go-profiles/internal/profile"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the profiles resource
func InitRouter(cfg *config.Configuration) *gin.Engine {
	router := gin.Default()
	db, _ := database.New(&cfg.DB)
	repo := profile.NewUserProfileRepository(db.Database)
	service := profile.NewUserProfileService(repo)

	userProfileController := profile.NewUserProfileController(service)

	MapRoutes(router, userProfileController)

	return router
}

// MapRoutes registers the routes of the resource
func MapRoutes(r *gin.Engine, userProfileController profile.UserProfileController) {
	userProfile := r.Group("/user-profiles")
	{
		userProfile.GET("/:id", userProfileController.Get)
		userProfile.POST("/:id", userProfileController.Create)
		userProfile.DELETE("/:id", userProfileController.Delete)
		userProfile.PUT("/:id", userProfileController.Update)
	}
}
