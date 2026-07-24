package raw_db_test

import (
	"testing"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql/raw_db"
)

func TestRawDatabase(t *testing.T) {
	t.Parallel()

	var _ database.RawDatabase = &raw_db.RawDatabase{}
}
