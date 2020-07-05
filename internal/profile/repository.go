package profile

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/prometheus/common/log"

	"github.com/CienciaArgentina/go-profiles/domain"
)

type userProfileRepository struct {
	db *sqlx.DB
}

// NewUserProfileRepository creates a new repository
func NewUserProfileRepository(db *sqlx.DB) UserProfileRepository {
	return &userProfileRepository{db}
}

func (u *userProfileRepository) Get(id int) (domain.UserProfile, error) {
	up := domain.UserProfile{}

	stm, err := u.db.PrepareNamed("SELECT * FROM UserProfile WHERE id = :id")
	if err != nil {
		log.Error(err)
		return up, fmt.Errorf("Unexpected error creating statement %v: %w", id, err)
	}

	err = stm.Get(&up, map[string]interface{}{"id": id})
	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return up, fmt.Errorf("UserProfile not found %v: %w", id, err)
		}
		return up, fmt.Errorf("Unexpected error getting UserProfile %v: %w", id, err)
	}

	return up, nil
}

func (u *userProfileRepository) Create(domain.UserProfile) error {
	return nil
}

func (u *userProfileRepository) Update(domain.UserProfile) error {
	return nil
}

func (u *userProfileRepository) Delete(int) error {
	return nil
}
