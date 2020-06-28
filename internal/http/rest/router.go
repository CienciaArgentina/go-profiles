package rest

import (
	"log"

	"github.com/CienciaArgentina/go-profiles/config"
	"github.com/CienciaArgentina/go-profiles/internal/profile"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the profiles resource
func InitRouter(cfg *config.Configuration) *gin.Engine {
	router := gin.Default()
	MapRoutes(router, cfg)
	return router
}

// MapRoutes registers the routes of the resource
func MapRoutes(r *gin.Engine, cfg *config.Configuration) {
	userProfileController, err := profile.NewUserProfileController(*cfg)

	if err != nil {
		log.Fatal(err)
	}

	userProfile := r.Group("/userProfiles")
	{
		userProfile.GET("/:id", userProfileController.Get)
		userProfile.POST("/:id", userProfileController.Create)
		userProfile.DELETE("/:id", userProfileController.Delete)
		userProfile.PUT("/:id", userProfileController.Update)

	}
}
