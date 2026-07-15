package postgresql

import (
	"testing"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql/user_db"
)

func TestUserDatabaseInterface(t *testing.T) {
	var _ database.UserDatabase = &user_db.UserDatabase{}
}
