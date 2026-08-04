package map_based

import (
	"crypto/rand"
	"errors"
	"sync"
	"time"

	"github.com/univers106/ITI/middlewares/sessions"
)

var (
	ErrCantGenerateKey = errors.New("could not generate key")
	ErrSessionNotFound = errors.New("session not found")
)

type SessionData struct {
	UserLogin string
	CreatedAt time.Time
	Timeout   time.Time
	LastVisit time.Time
}

type MapBasedSessionStorage struct {
	sessions map[string]SessionData
	mu       sync.Mutex
}

func (m *MapBasedSessionStorage) GetLoginFromSession(key string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	data, ok := m.sessions[key]
	if !ok {
		return "", ErrSessionNotFound
	}

	now := time.Now()
	if now.After(data.Timeout) || now.After(data.LastVisit.Add(sessions.SessionIdleTimeout)) {
		delete(m.sessions, key)

		return "", ErrSessionNotFound
	}

	data.LastVisit = time.Now()
	m.sessions[key] = data

	return data.UserLogin, nil
}

func (m *MapBasedSessionStorage) NewSession(userLogin string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	const maxAttempt = 20

	var key string
	for range maxAttempt {
		key = rand.Text()
		if _, ok := m.sessions[key]; !ok {
			now := time.Now()
			m.sessions[key] = SessionData{
				UserLogin: userLogin,
				CreatedAt: now,
				Timeout:   now.Add(sessions.SessionTimeout),
				LastVisit: now,
			}

			return key, nil
		}
	}

	return "", ErrCantGenerateKey
}

func (m *MapBasedSessionStorage) DeleteSession(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.sessions[key]
	if !ok {
		return ErrSessionNotFound
	}

	delete(m.sessions, key)

	return nil
}

func (m *MapBasedSessionStorage) DeleteAllUserSessions(userLogin string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for key, data := range m.sessions {
		if data.UserLogin == userLogin {
			delete(m.sessions, key)
		}
	}

	return nil
}

func NewSessionStorage() *MapBasedSessionStorage {
	return &MapBasedSessionStorage{
		sessions: make(map[string]SessionData),
	}
}
