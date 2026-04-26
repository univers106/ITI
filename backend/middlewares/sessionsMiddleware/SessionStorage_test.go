package sessionsMiddleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	sessionsMiddleware "github.com/univers106/ITI/middlewares/sessionsMiddleware"
)

func TestNewSessionStorage(t *testing.T) {
	t.Parallel()

	storage := sessionsMiddleware.NewSessionStorage()
	assert.NotNil(t, storage)
}

func TestSessionStorage_NewSession(t *testing.T) {
	t.Parallel()

	storage := sessionsMiddleware.NewSessionStorage()
	userId := 42

	key, err := storage.NewSession(userId)
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	retrievedId, err := storage.GetIdFromSession(key)
	require.NoError(t, err)
	assert.Equal(t, userId, retrievedId)
}

func TestSessionStorage_GetIdFromSession_NotFound(t *testing.T) {
	t.Parallel()

	storage := sessionsMiddleware.NewSessionStorage()

	_, err := storage.GetIdFromSession("nonexistent")
	require.Error(t, err)
	assert.Equal(t, sessionsMiddleware.ErrSessionNotFound, err)
}

func TestSessionStorage_DeleteSession(t *testing.T) {
	t.Parallel()

	storage := sessionsMiddleware.NewSessionStorage()
	userId := 42
	key, _ := storage.NewSession(userId)

	err := storage.DeleteSession(key)
	require.NoError(t, err)

	_, err = storage.GetIdFromSession(key)
	require.Error(t, err)
	assert.Equal(t, sessionsMiddleware.ErrSessionNotFound, err)
}

func TestSessionStorage_DeleteSession_NotFound(t *testing.T) {
	t.Parallel()

	storage := sessionsMiddleware.NewSessionStorage()

	err := storage.DeleteSession("nonexistent")
	require.Error(t, err)
	assert.Equal(t, sessionsMiddleware.ErrSessionNotFound, err)
}

func TestSessionStorage_DeleteUserSessions(t *testing.T) {
	t.Parallel()

	storage := sessionsMiddleware.NewSessionStorage()
	userId1 := 1
	userId2 := 2

	key1, _ := storage.NewSession(userId1)
	key2, _ := storage.NewSession(userId1)
	key3, _ := storage.NewSession(userId2)

	err := storage.DeleteUserSessions(userId1)
	require.NoError(t, err)

	_, err = storage.GetIdFromSession(key1)
	require.Error(t, err)
	_, err = storage.GetIdFromSession(key2)
	require.Error(t, err)

	id, err := storage.GetIdFromSession(key3)
	require.NoError(t, err)
	assert.Equal(t, userId2, id)
}
