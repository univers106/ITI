package database_middleware_test

import (
	"errors"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/univers106/ITI/database"
	dbMiddleware "github.com/univers106/ITI/middlewares/database_middleware"
)

var errNotImplemented = errors.New("mock: not implemented")

type mockDatabase struct{}

func (m *mockDatabase) GetUserByID(id int) (*database.User, error) {
	return nil, errNotImplemented
}

func (m *mockDatabase) GetUserByLogin(login string) (*database.User, error) {
	return nil, errNotImplemented
}

func (m *mockDatabase) UserAuthentication(login string, password string) (*database.User, error) {
	return nil, errNotImplemented
}

func (m *mockDatabase) CreateUser(login string, name string, password string) error {
	return errNotImplemented
}

func (m *mockDatabase) DeleteUser(id int) error {
	return errNotImplemented
}

func (m *mockDatabase) ChangeUserPassword(user_id int, password string) error {
	return errNotImplemented
}

func (m *mockDatabase) ChangeUserLogin(user_id int, login string) error {
	return errNotImplemented
}

func (m *mockDatabase) ChangeUserName(user_id int, name string) error {
	return errNotImplemented
}

func (m *mockDatabase) UserAddPermissions(user_id int, permission string) error {
	return errNotImplemented
}

func (m *mockDatabase) UserRemovePermissions(user_id int, permission string) error {
	return errNotImplemented
}

func (m *mockDatabase) UserCheckPermission(user_id int, permission string) (bool, error) {
	return false, errNotImplemented
}

func TestNewDatabaseMiddleware(t *testing.T) {
	t.Parallel()

	db := &mockDatabase{}
	middleware := dbMiddleware.NewDatabaseMiddleware(db)

	e := echo.New()
	c := e.NewContext(nil, nil)

	handler := middleware(func(c *echo.Context) error {
		retrievedDB, err := dbMiddleware.GetDatabase(c)
		require.NoError(t, err)
		assert.Equal(t, db, retrievedDB)

		return nil
	})

	err := handler(c)
	require.NoError(t, err)
}

func TestGetDatabase_NotFound(t *testing.T) {
	t.Parallel()

	e := echo.New()
	c := e.NewContext(nil, nil)

	_, err := dbMiddleware.GetDatabase(c)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get session store")
}

func TestGetDatabase_Success(t *testing.T) {
	t.Parallel()

	db := &mockDatabase{}
	e := echo.New()
	c := e.NewContext(nil, nil)
	c.Set("_database", db)

	retrievedDB, err := dbMiddleware.GetDatabase(c)
	require.NoError(t, err)
	assert.Equal(t, db, retrievedDB)
}
