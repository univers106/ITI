package cli

import (
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql/user_db"
)

// Run runs the CLI
// returns true if has command.
func Run(pool *pgxpool.Pool) bool {
	userDB := user_db.NewUserDatabase(pool)

	argsWithProg := os.Args

	if len(argsWithProg) <= 1 {
		return false
	}

	if argsWithProg[1] == "create-super-user" {
		if len(argsWithProg) < 4 { //nolint
			slog.Error("использование: create-super-user [login] [password]")

			return true
		}

		user := database.User{
			Login:       argsWithProg[2],
			Name:        argsWithProg[2],
			Permissions: []string{database.PermSuperUser},
		}
		user.SetPassword(argsWithProg[3])

		err := userDB.Create(user)
		if err != nil {
			panic("failed to create super user: " + err.Error())
		}

		slog.Warn("Создан супер пользователь, убедитесь что удалите его после использования")

		return true
	}

	return true
}
