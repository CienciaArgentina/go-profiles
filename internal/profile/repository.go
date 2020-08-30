package profile

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/CienciaArgentina/go-profiles/domain"
	"github.com/CienciaArgentina/go-profiles/internal/errors"

	log "github.com/sirupsen/logrus"
)

const (
	dbName       = "profile"
	profilesName = "profiles"
)

type userProfileRepository struct {
	db *mongo.Client
}

// NewUserProfileRepository creates a new repository
func NewUserProfileRepository(db *mongo.Client) UserProfileRepository {
	return &userProfileRepository{db}
}

func (u *userProfileRepository) Get(id int) (domain.UserProfile, error) {
	profiles := profiles(*u.db)

	result := domain.UserProfile{}

	filter := bson.D{primitive.E{Key: "userid", Value: id}}

	queryResult := profiles.FindOne(context.TODO(), filter)

	if err := queryResult.Err(); err != nil {
		log.Error("Error getting UserProfile: ", err)

		if err == mongo.ErrNoDocuments {
			return result, errors.ErrUserProfileNotFound
		}

		return result, errors.ErrInternalServerError
	}

	if err := queryResult.Decode(&result); err != nil {
		log.Error("Error getting UserProfile: ", err)
		return result, errors.ErrInternalServerError
	}

	return result, nil
}

func (u *userProfileRepository) Create(userProfile domain.UserProfile) error {
	profiles := profiles(*u.db)

	result, err := profiles.InsertOne(context.TODO(), userProfile)

	if err != nil {
		log.Error("Error creating UserProfile: ", err)
		return errors.ErrInternalServerError
	}

	log.Debug(fmt.Sprintf("Created UserProfile with ID %d", result.InsertedID))

	return nil
}

func (u *userProfileRepository) Update(userProfile domain.UserProfile) error {
	profiles := profiles(*u.db)

	ctx := context.TODO()
	filter := bson.D{primitive.E{Key: "userid", Value: userProfile.UserID}}
	document := bson.M{"$set": userProfile}

	result, err := profiles.UpdateOne(ctx, filter, document)

	if err != nil {
		log.Error("Error updating UserProfile", err)
	}

	if result.MatchedCount == 0 {
		log.Warn("UserProfile not found", filter)
		return errors.ErrUserProfileNotFound
	}

	return nil
}

func (u *userProfileRepository) Delete(id int) error {
	profiles := profiles(*u.db)

	filter := bson.D{primitive.E{Key: "userid", Value: id}}

	result, err := profiles.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Error("Error deleting UserProfile", err)
		return errors.ErrInternalServerError
	}

	if result.DeletedCount == 0 {
		log.Warn("UserProfile not found", filter)
		return errors.ErrUserProfileNotFound
	}

	return nil
}

func profiles(client mongo.Client) *mongo.Collection {
	return client.Database(dbName).Collection(profilesName)
}
