package sessions_middleware_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/middlewares/sessions_middleware"
)

func TestNewSessionsMiddleware(t *testing.T) {
	t.Parallel()

	store := sessions_middleware.NewSessionStorage()
	middleware := sessions_middleware.NewSessionsMiddleware(store)

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := middleware(func(c *echo.Context) error {
		retrievedStore, err := sessions_middleware.GetSessionStorage(c)
		require.NoError(t, err)
		assert.Equal(t, store, retrievedStore)

		return nil
	})

	err := handler(c)
	require.NoError(t, err)
}

func TestGetSessionStorage_NotFound(t *testing.T) {
	t.Parallel()

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_, err := sessions_middleware.GetSessionStorage(c)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get session store")
}

func TestGetKeyFromCookies_NoCookie(t *testing.T) {
	t.Parallel()

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_, err := sessions_middleware.GetKeyFromCookies(c)
	require.Error(t, err)
	assert.Equal(t, sessions_middleware.ErrUnauthorized, err)
}

func TestGetKeyFromCookies_WithCookie(t *testing.T) {
	t.Parallel()

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "session_key", Value: "abc123"})

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	key, err := sessions_middleware.GetKeyFromCookies(c)
	require.NoError(t, err)
	assert.Equal(t, "abc123", key)
}

func TestSetKeyToCookies(t *testing.T) {
	t.Parallel()

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	sessions_middleware.SetKeyToCookies(c, "testkey")

	cookies := rec.Result().Cookies()
	assert.Len(t, cookies, 1)
	cookie := cookies[0]
	assert.Equal(t, "session_key", cookie.Name)
	assert.Equal(t, "testkey", cookie.Value)
	assert.True(t, cookie.HttpOnly)
	assert.True(t, cookie.Secure)
	assert.Equal(t, http.SameSiteStrictMode, cookie.SameSite)
	assert.Positive(t, cookie.MaxAge)
}

func TestDeleteCookies(t *testing.T) {
	t.Parallel()

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	sessions_middleware.DeleteCookies(c)

	cookies := rec.Result().Cookies()
	assert.Len(t, cookies, 1)
	cookie := cookies[0]
	assert.Equal(t, "session_key", cookie.Name)
	assert.Empty(t, cookie.Value)
	assert.Equal(t, -1, cookie.MaxAge)
	assert.True(t, cookie.Expires.Before(time.Now()))
}

func TestGetUserFromSession_Success(t *testing.T) {
	t.Parallel()

	store := sessions_middleware.NewSessionStorage()
	db := &mockDatabase{
		users: map[int]*database.User{
			1: {ID: 1, Login: "user1", Name: "User One"},
		},
	}

	key, err := store.NewSession(1)
	require.NoError(t, err)

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "session_key", Value: key})

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("_session_storage", store)
	c.Set("_database", db)

	user, err := sessions_middleware.GetUserFromSession(c)
	require.NoError(t, err)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "user1", user.Login)
}

func TestGetUserFromSession_NoCookie(t *testing.T) {
	t.Parallel()

	store := sessions_middleware.NewSessionStorage()
	db := &mockDatabase{}

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("_session_storage", store)
	c.Set("_database", db)

	_, err := sessions_middleware.GetUserFromSession(c)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get user from session")
}

func TestGetUserFromSession_SessionNotFound(t *testing.T) {
	t.Parallel()

	store := sessions_middleware.NewSessionStorage()
	db := &mockDatabase{}

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "session_key", Value: "invalidkey"})

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("_session_storage", store)
	c.Set("_database", db)

	_, err := sessions_middleware.GetUserFromSession(c)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get user from session")
}

func TestGetUserFromSession_DatabaseError(t *testing.T) {
	t.Parallel()

	store := sessions_middleware.NewSessionStorage()
	db := &mockDatabase{
		users: map[int]*database.User{},
	}

	key, err := store.NewSession(999)
	require.NoError(t, err)

	echoServer := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "session_key", Value: key})

	rec := httptest.NewRecorder()
	context := echoServer.NewContext(req, rec)

	context.Set("_session_storage", store)
	context.Set("_database", db)

	_, err = sessions_middleware.GetUserFromSession(context)
	require.Error(t, err)

	var httpErr *echo.HTTPError

	ok := errors.As(err, &httpErr)
	assert.True(t, ok)
	assert.Equal(t, http.StatusInternalServerError, httpErr.Code)
	assert.Contains(t, httpErr.Message, "can't get user by login")
}

func TestGetUserFromSession_NoDatabaseInContext(t *testing.T) {
	t.Parallel()

	store := sessions_middleware.NewSessionStorage()
	key, _ := store.NewSession(1)

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "session_key", Value: key})

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("_session_storage", store)

	_, err := sessions_middleware.GetUserFromSession(c)
	require.Error(t, err)

	var httpErr *echo.HTTPError

	ok := errors.As(err, &httpErr)
	assert.True(t, ok)
	assert.Equal(t, http.StatusInternalServerError, httpErr.Code)
	assert.Contains(t, httpErr.Message, "failed to get database")
}

func TestGetUserFromSession_NoSessionStorageInContext(t *testing.T) {
	t.Parallel()

	e := echo.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_, err := sessions_middleware.GetUserFromSession(c)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get user from session")
}
