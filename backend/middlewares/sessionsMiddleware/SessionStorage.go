package sessionsMiddleware

import (
	"crypto/rand"
	"errors"
	"sync"
	"time"
)

var (
	ErrCantGenerateKey = errors.New("could not generate key")
	ErrSessionNotFound = errors.New("session not found")
)

const (
	SessionTimeout     = 3 * time.Hour
	SessionIdleTimeout = 10 * time.Minute
)

type SessionStorage interface {
	GetIdFromSession(key string) (int, error)
	NewSession(userId int) (string, error)
	DeleteSession(key string) error
	DeleteUserSessions(userId int) error
}

// далее реализация на map, если, у вас сервис больше,
// то стоит сделать реализацию на субд

type SessionData struct {
	UserId    int
	CreatedAt time.Time
	Timeout   time.Time
	LastVisit time.Time
}

type MapBasedSessionStorage struct {
	sessions map[string]SessionData
	mu       sync.Mutex
}

func (m *MapBasedSessionStorage) GetIdFromSession(key string) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	data, ok := m.sessions[key]
	if !ok {
		return 0, ErrSessionNotFound
	}

	now := time.Now()
	if now.After(data.Timeout) || now.After(data.LastVisit.Add(SessionIdleTimeout)) {
		delete(m.sessions, key)

		return 0, ErrSessionNotFound
	}

	data.LastVisit = time.Now()
	m.sessions[key] = data

	return data.UserId, nil
}

func (m *MapBasedSessionStorage) NewSession(userId int) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	const maxAttempt = 20

	var key string
	for range maxAttempt {
		key = rand.Text()
		if _, ok := m.sessions[key]; !ok {
			now := time.Now()
			m.sessions[key] = SessionData{
				UserId:    userId,
				CreatedAt: now,
				Timeout:   now.Add(SessionTimeout),
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

func (m *MapBasedSessionStorage) DeleteUserSessions(userId int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for key, data := range m.sessions {
		if data.UserId == userId {
			delete(m.sessions, key)
		}
	}

	return nil
}

func NewSessionStorage() SessionStorage {
	return &MapBasedSessionStorage{
		sessions: make(map[string]SessionData),
	}
}
