package profile

import (
	"errors"
	"testing"

	"github.com/CienciaArgentina/go-profiles/domain"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	profiles := []domain.UserProfile{
		{UserID: 1, UserName: "cosme.fulanito", Name: "Cosme", LastName: "Fulanito", Email: "cosme@fulanito.com"},
		{UserID: 2, UserName: "john.doe", Name: "John", LastName: "Doe", Email: "john@doe.com"},
	}

	t.Run("get existent user", func(t *testing.T) {
		assert := assert.New(t)

		db, mock := mockDB(t)
		defer db.DB.Close()

		mock.ExpectPrepare("^SELECT (.+) FROM UserProfile WHERE").
			ExpectQuery().
			WithArgs(1).
			WillReturnRows(mockRows(profiles[:1]))

		repo := NewUserProfileRepository(db)

		p, err := repo.Get(1)

		assert.NoError(err)
		assert.NotEmpty(p)
		assert.Equal(profiles[0], p)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})

	t.Run("get non existent user", func(t *testing.T) {
		assert := assert.New(t)

		db, mock := mockDB(t)
		defer db.DB.Close()

		mock.ExpectPrepare("^SELECT (.+) FROM UserProfile WHERE").
			ExpectQuery().
			WithArgs(2).
			WillReturnRows(mockRows([]domain.UserProfile{}))

		repo := NewUserProfileRepository(db)

		p, err := repo.Get(2)

		assert.Error(err)
		assert.Regexp("^UserProfile not found", err.Error())
		assert.Empty(p)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})
	t.Run("database problem", func(t *testing.T) {
		assert := assert.New(t)

		db, mock := mockDB(t)
		defer db.DB.Close()

		mock.ExpectPrepare("^SELECT (.+) FROM UserProfile WHERE").
			ExpectQuery().
			WillReturnError(errors.New("some error"))

		repo := NewUserProfileRepository(db)

		p, err := repo.Get(2)

		assert.Error(err)
		assert.Regexp("^Unexpected error getting", err.Error())
		assert.Empty(p)
	})
}

func mockDB(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	return sqlxDB, mock
}

func mockRows(data []domain.UserProfile) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "username", "name", "last_name", "email"})

	for _, up := range data {
		rows = rows.AddRow(up.UserID, up.UserName, up.Name, up.LastName, up.Email)
	}

	return rows

}
