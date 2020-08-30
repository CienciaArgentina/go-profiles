package profile

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/CienciaArgentina/go-profiles/domain"
	"github.com/CienciaArgentina/go-profiles/internal/errors"
)

var (
	john = domain.UserProfile{
		UserID:      1,
		Name:        "John",
		LastName:    "Doe",
		Email:       "a@b.c",
		Address:     domain.Address{City: "bla"},
		Birthday:    time.Now(),
		UserBlocked: false,
	}
	newUser          = domain.UserProfile{UserID: 99}
	newUserWithError = domain.UserProfile{UserID: -1}
	errUnknown       = fmt.Errorf("Unknown")
)

type mockUserProfileRepo struct {
	profiles []domain.UserProfile
}

func (u *mockUserProfileRepo) Get(id int) (domain.UserProfile, error) {
	if id == -1 {
		return domain.UserProfile{}, errUnknown
	}

	for i := range u.profiles {
		if u.profiles[i].UserID == id {
			return u.profiles[i], nil
		}
	}
	return domain.UserProfile{}, errors.ErrUserProfileNotFound

}
func (u *mockUserProfileRepo) Create(up domain.UserProfile) error {
	if up.UserID == -1 {
		return errUnknown
	}

	u.profiles = append(u.profiles, up)

	return nil
}
func (u *mockUserProfileRepo) Update(up domain.UserProfile) error {
	if up.UserID == -1 {
		return errUnknown
	}

	for i := range u.profiles {
		if u.profiles[i].UserID == up.UserID {
			u.profiles[i] = up
			return nil
		}
	}
	return errors.ErrUserProfileNotFound
}

func (u *mockUserProfileRepo) Delete(id int) error {
	if id == -1 {
		return errUnknown
	}

	for i := range u.profiles {
		if u.profiles[i].UserID == id {
			u.profiles[i] = u.profiles[len(u.profiles)-1]
			u.profiles = u.profiles[:len(u.profiles)-1]
			return nil
		}
	}
	return errors.ErrUserProfileNotFound
}

func TestNewUserProfileService(t *testing.T) {
	t.Run("create valid service", func(t *testing.T) {
		if got := NewUserProfileService(&mockUserProfileRepo{}); got == nil {
			t.Errorf("expected a valid services, got nil")
		}
	})
}

func Test_userProfileService_FindUserProfile(t *testing.T) {
	mockService := mockService()

	tests := []struct {
		name    string
		id      int
		want    domain.UserProfile
		wantErr error
	}{
		{"query valid user", 1, john, nil},
		{"query invalid user", 2, domain.UserProfile{}, errors.ErrUserProfileNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mockService.FindUserProfile(tt.id)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("userProfileService.FindUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userProfileService.FindUserProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userProfileService_CreateUserProfile(t *testing.T) {
	mockService := mockService()
	tests := []struct {
		name    string
		user    domain.UserProfile
		wantErr error
	}{
		{"create valid user", newUser, nil},
		{"create valid user but fails", newUserWithError, errUnknown},
		{"create existent user", john, errors.ErrUserProfileAlreadyExists},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mockService.CreateUserProfile(tt.user)
			if err != tt.wantErr {
				t.Errorf("userProfileService.FindUserProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userProfileService_UpdateUserProfile(t *testing.T) {
	mockService := mockService()
	type args struct {
		id   int
		user domain.UserProfile
	}
	tests := []struct {
		name string
		args
		wantErr error
	}{
		{"update valid user", args{john.UserID, newUser}, nil},
		{"update valid user but fails", args{newUserWithError.UserID, newUserWithError}, errUnknown},
		{"update invalid user", args{99, newUser}, errors.ErrUserProfileNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mockService.UpdateUserProfile(tt.args.id, tt.args.user)
			if err != tt.wantErr {
				t.Errorf("userProfileService.FindUserProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userProfileService_DeleteUserProfile(t *testing.T) {
	mockService := mockService()
	tests := []struct {
		name    string
		id      int
		wantErr error
	}{
		{"delete valid user", 1, nil},
		{"delete invalid user", 999, errors.ErrUserProfileNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mockService.DeleteUserProfile(tt.id)
			if err != tt.wantErr {
				t.Errorf("userProfileService.FindUserProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func mockService() UserProfileService {
	repo := mockUserProfileRepo{[]domain.UserProfile{john}}
	return NewUserProfileService(&repo)
}
