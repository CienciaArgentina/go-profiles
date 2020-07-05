package profile

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

type userProfileController struct {
	service UserProfileService
}

// NewUserProfileController creates a new controller
func NewUserProfileController(service UserProfileService) UserProfileController {
	return &userProfileController{service}
}

// Get gets a UserProfile by id
func (c *userProfileController) Get(ctx *gin.Context) {
	rawID := ctx.Params.ByName("id")
	id, err := strconv.Atoi(rawID)

	if err != nil {
		err = ctx.AbortWithError(http.StatusBadRequest,
			apierror.NewWithStatus(http.StatusBadRequest).WithMessage(fmt.Sprintf("%s: invalid format", rawID)))
		if err != nil {
			log.Error("Error while aborting", err)
		}
	} else {

		up, err := c.service.FindUserProfile(id)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			ctx.JSON(http.StatusOK, up)
		}
	}
}

func (c *userProfileController) Create(ctx *gin.Context) {
	// @TODO implement me
	ctx.Status(http.StatusOK)
}

func (c *userProfileController) Delete(ctx *gin.Context) {
	// @TODO implement me
	ctx.Status(http.StatusOK)
}

func (c *userProfileController) Update(ctx *gin.Context) {
	// @TODO implement me
	ctx.Status(http.StatusOK)
}
