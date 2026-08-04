package sessions_test

import (
	"testing"

	"github.com/univers106/ITI/middlewares/sessions"
	"github.com/univers106/ITI/middlewares/sessions/map_based"
)

func TestSessionStorageInterface(t *testing.T) {
	var _ sessions.SessionStorage = map_based.NewSessionStorage()

}
