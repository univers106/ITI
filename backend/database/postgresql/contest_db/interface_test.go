package contest_db_test

import (
	"testing"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql/contest_db"
)

func TestContestDatabase(t *testing.T) {
	t.Parallel()

	var _ database.ContestDatabase = &contest_db.ContestDatabase{}
}
