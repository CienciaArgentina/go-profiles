package profile

import (
	"net/http"

	"github.com/CienciaArgentina/go-profiles/config"
	"github.com/CienciaArgentina/go-profiles/domain"
	"github.com/gin-gonic/gin"
)

type userProfileController struct {
}

// NewUserProfileController creates a new controller
func NewUserProfileController(config.Configuration) (UserProfileController, error) {
	// @TODO implement me
	return new(userProfileController), nil
}

// Get bla
func (c userProfileController) Get(ctx *gin.Context) {
	// @TODO implement me
	ctx.JSON(http.StatusOK, new(domain.UserProfile))
}

func (c userProfileController) Create(ctx *gin.Context) {
	// @TODO implement me
	ctx.Status(http.StatusOK)
}

func (c userProfileController) Delete(ctx *gin.Context) {
	// @TODO implement me
	ctx.Status(http.StatusOK)
}

func (c userProfileController) Update(ctx *gin.Context) {
	// @TODO implement me
	ctx.Status(http.StatusOK)
}
