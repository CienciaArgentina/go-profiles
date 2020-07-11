package rest

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CienciaArgentina/go-profiles/config"
	"github.com/CienciaArgentina/go-profiles/internal/profile"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the profiles resource
func InitRouter(cfg *config.Configuration) (*gin.Engine, func()) {
	router := gin.Default()
	db := NewMongoClient(cfg)
	repo := profile.NewUserProfileRepository(db)
	service := profile.NewUserProfileService(repo)

	userProfileController := profile.NewUserProfileController(service)

	MapRoutes(router, userProfileController)

	return router, func() { db.Disconnect(context.TODO()) }
}

// MapRoutes registers the routes of the resource
func MapRoutes(r *gin.Engine, userProfileController profile.UserProfileController) {
	userProfile := r.Group("/user-profiles")
	{
		userProfile.GET("/:id", userProfileController.Get)
		userProfile.POST("/", userProfileController.Create)
		userProfile.DELETE("/:id", userProfileController.Delete)
		userProfile.PUT("/:id", userProfileController.Update)
	}
}

// NewMongoClient creates a new client to mongo db
func NewMongoClient(cfg *config.Configuration) *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		cfg.DB.Username, cfg.DB.Password, cfg.DB.Hostname, cfg.DB.Port)

	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
