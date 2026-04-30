package sessions_middleware_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

var errNotImplemented = errors.New("mock: not implemented")

type mockDatabase struct {
	users map[int]*database.User
}

func (m *mockDatabase) GetUserByID(id int) (*database.User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, database.ErrUserNotFound
	}

	return user, nil
}

func (m *mockDatabase) GetUserByLogin(login string) (*database.User, error) {
	for _, user := range m.users {
		if user.Login == login {
			return user, nil
		}
	}

	return nil, database.ErrUserNotFound
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

func TestOnlyUsersMiddleware_Success(t *testing.T) {
	t.Parallel()

	store := sessions_middleware.NewSessionStorage()
	userId := 1
	key, err := store.NewSession(userId)
	require.NoError(t, err)

	db := &mockDatabase{
		users: map[int]*database.User{
			1: {ID: 1, Login: "user1", Name: "User One"},
		},
	}

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "session_key", Value: key})

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("_session_storage", store)
	c.Set("_database", db)

	var capturedUser *database.User

	handler := sessions_middleware.OnlyUsersMiddleware(func(c *echo.Context) error {
		user, err := sessions_middleware.GetUser(c)
		require.NoError(t, err)
		assert.Equal(t, userId, user.ID)
		capturedUser = user

		return nil
	})

	err = handler(c)
	require.NoError(t, err)
	assert.NotNil(t, capturedUser)
	assert.Equal(t, userId, capturedUser.ID)
}

func TestOnlyUsersMiddleware_NoCookie(t *testing.T) {
	t.Parallel()

	store := sessions_middleware.NewSessionStorage()
	db := &mockDatabase{}

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("_session_storage", store)
	c.Set("_database", db)

	handler := sessions_middleware.OnlyUsersMiddleware(func(c *echo.Context) error {
		t.Fatal("should not be called")

		return nil
	})

	err := handler(c)
	require.Error(t, err)

	var httpErr *echo.HTTPError

	ok := errors.As(err, &httpErr)
	assert.True(t, ok)
	assert.Equal(t, http.StatusUnauthorized, httpErr.Code)
	assert.Contains(t, httpErr.Message, "unauthorized")
}

func TestOnlyUsersMiddleware_SessionNotFound(t *testing.T) {
	t.Parallel()

	store := sessions_middleware.NewSessionStorage()
	db := &mockDatabase{}

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "session_key", Value: "invalid"})

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("_session_storage", store)
	c.Set("_database", db)

	handler := sessions_middleware.OnlyUsersMiddleware(func(c *echo.Context) error {
		t.Fatal("should not be called")

		return nil
	})

	err := handler(c)
	require.Error(t, err)
}

func TestOnlyUsersMiddleware_DatabaseError(t *testing.T) {
	t.Parallel()

	store := sessions_middleware.NewSessionStorage()
	key, _ := store.NewSession(999)
	db := &mockDatabase{users: map[int]*database.User{}}

	echoServer := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)

	req.AddCookie(&http.Cookie{Name: "session_key", Value: key})

	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)

	c.Set("_session_storage", store)
	c.Set("_database", db)

	handler := sessions_middleware.OnlyUsersMiddleware(func(c *echo.Context) error {
		t.Fatal("should not be called")

		return nil
	})

	err := handler(c)
	require.Error(t, err)

	var httpErr *echo.HTTPError

	ok := errors.As(err, &httpErr)
	assert.True(t, ok)
	assert.Equal(t, http.StatusInternalServerError, httpErr.Code)
}

func TestGetUser_Success(t *testing.T) {
	t.Parallel()

	e := echo.New()
	c := e.NewContext(nil, nil)
	expectedUser := &database.User{ID: 5, Login: "test"}
	c.Set("user", expectedUser)

	user, err := sessions_middleware.GetUser(c)
	require.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestGetUser_NotFound(t *testing.T) {
	t.Parallel()

	e := echo.New()
	c := e.NewContext(nil, nil)

	_, err := sessions_middleware.GetUser(c)
	require.Error(t, err)
}

func TestGetUserDbCheckPermision_Success(t *testing.T) {
	t.Parallel()

	e := echo.New()
	c := e.NewContext(nil, nil)
	user := &database.User{
		ID:          7,
		Permissions: []string{"UsersManipulation"},
	}
	c.Set("user", user)

	db := &mockDatabase{}
	c.Set("_database", db)

	retrievedUser, retrievedDb, httpErr := sessions_middleware.GetUserDbCheckPermision(
		c,
		"UsersManipulation",
	)
	assert.Nil(t, httpErr)
	assert.Equal(t, user, retrievedUser)
	assert.Equal(t, db, retrievedDb)
}

func TestGetUserDbCheckPermision_NoUserInContext(t *testing.T) {
	t.Parallel()

	e := echo.New()
	c := e.NewContext(nil, nil)
	db := &mockDatabase{}
	c.Set("_database", db)

	_, _, httpErr := sessions_middleware.GetUserDbCheckPermision(c, "UsersManipulation")
	assert.NotNil(t, httpErr)
	assert.Equal(t, http.StatusInternalServerError, httpErr.Code)
	assert.Contains(t, httpErr.Message, "failed to get user from context")
}

func TestGetUserDbCheckPermision_NoPermission(t *testing.T) {
	t.Parallel()

	e := echo.New()
	c := e.NewContext(nil, nil)
	user := &database.User{
		ID:          7,
		Permissions: []string{"OtherPermission"},
	}
	c.Set("user", user)

	db := &mockDatabase{}
	c.Set("_database", db)

	_, _, httpErr := sessions_middleware.GetUserDbCheckPermision(c, "UsersManipulation")
	assert.NotNil(t, httpErr)
	assert.Equal(t, http.StatusForbidden, httpErr.Code)
	assert.Contains(t, httpErr.Message, "You do not have permission")
}

func TestGetUserDbCheckPermision_SuperUserBypass(t *testing.T) {
	t.Parallel()

	e := echo.New()
	c := e.NewContext(nil, nil)

	user := &database.User{
		ID:          7,
		Permissions: []string{database.PermSuperUser},
	}

	c.Set("user", user)

	db := &mockDatabase{}
	c.Set("_database", db)

	retrievedUser, retrievedDb, httpErr := sessions_middleware.GetUserDbCheckPermision(
		c,
		"UsersManipulation",
	)
	assert.Nil(t, httpErr)
	assert.Equal(t, user, retrievedUser)
	assert.Equal(t, db, retrievedDb)
}

func TestGetUserDbCheckPermision_NoDatabaseInContext(t *testing.T) {
	t.Parallel()

	e := echo.New()
	c := e.NewContext(nil, nil)
	user := &database.User{
		ID:          7,
		Permissions: []string{"UsersManipulation"},
	}
	c.Set("user", user)

	_, _, httpErr := sessions_middleware.GetUserDbCheckPermision(c, "UsersManipulation")
	assert.NotNil(t, httpErr)
	assert.Equal(t, http.StatusInternalServerError, httpErr.Code)
	assert.Contains(t, httpErr.Message, "failed to get database")
}
