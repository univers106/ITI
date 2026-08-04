package sessions

type SessionStorage interface {
	GetLoginFromSession(key string) (string, error)
	NewSession(userLogin string) (string, error)
	DeleteAllUserSessions(userLogin string) error
	DeleteSession(key string) error
}
