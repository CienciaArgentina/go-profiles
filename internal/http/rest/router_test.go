package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CienciaArgentina/go-profiles/domain"
	"github.com/CienciaArgentina/go-profiles/internal/errors"
	"github.com/CienciaArgentina/go-profiles/internal/profile"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type userProfileServiceMock struct {
	mock.Mock
}

func (u *userProfileServiceMock) FindUserProfile(id int) (domain.UserProfile, error) {
	args := u.Called(id)
	return args.Get(0).(domain.UserProfile), args.Error(1)
}

func (u *userProfileServiceMock) CreateUserProfile(domain.UserProfile) error {
	return nil
}

func (u *userProfileServiceMock) UpdateUserProfile(int, domain.UserProfile) error {
	return nil
}

func (u *userProfileServiceMock) DeleteUserProfile(int) error {
	return nil
}

func TestUserProfileController(t *testing.T) {
	var profileGot domain.UserProfile
	service := new(userProfileServiceMock)
	profileExpected := domain.UserProfile{
		UserID: 1, UserName: "john.dow", Name: "John", LastName: "Doe",
	}

	service.On("FindUserProfile", 1).Return(profileExpected, nil)
	service.On("FindUserProfile", 2).Return(domain.UserProfile{}, errors.ErrUserProfileNotFound)

	router := gin.Default()

	MapRoutes(router, profile.NewUserProfileController(service))

	t.Run("/get valid user", func(t *testing.T) {
		statusExpected := http.StatusOK
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/user-profiles/1", nil)

		router.ServeHTTP(w, r)

		_ = json.Unmarshal(w.Body.Bytes(), &profileGot)

		require.Equal(t, statusExpected, w.Code)
		require.Equal(t, profileExpected, profileGot)
	})

	t.Run("/get invalid user", func(t *testing.T) {
		statusExpected := http.StatusNotFound
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/user-profiles/2", nil)

		router.ServeHTTP(w, r)

		require.Equal(t, statusExpected, w.Code)
		require.Empty(t, w.Body.Bytes())
	})
}
