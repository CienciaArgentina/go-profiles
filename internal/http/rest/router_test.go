package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CienciaArgentina/go-profiles/config"
	"github.com/CienciaArgentina/go-profiles/domain"
	"github.com/stretchr/testify/require"
)

func TestUserProfileController(t *testing.T) {
	t.Run("/get valid user", func(t *testing.T) {
		statusWanted := http.StatusOK
		bodyWanted, _ := json.Marshal(new(domain.UserProfile))
		cfg := config.Configuration{}

		router := InitRouter(&cfg)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/userProfiles/1", nil)
		router.ServeHTTP(w, r)

		require.Equal(t, w.Code, statusWanted)
		require.Equal(t, w.Body.Bytes(), bodyWanted)
	})
}
