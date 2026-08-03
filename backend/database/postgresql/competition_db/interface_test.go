package competition_db_test

import (
	"testing"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql/competition_db"
)

func Test(t *testing.T) {
	var _ database.CompetitionDatabase = &competition_db.CompetitionDatabase{}
}
