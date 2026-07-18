package cli

import (
	"log/slog"
	"os"

	"github.com/univers106/ITI/config"
	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
	"github.com/univers106/ITI/database/postgresql/user_db"
)

func get_conf() config.Config {
	return config.ReadConfig("config.yaml")
}

// Run runs the CLI
// returns true if has command.
func Run() bool {
	cfg := get_conf()

	pgPool := postgresql.NewPool(cfg.PostgresSqlURL)
	userDB := user_db.NewUserDatabase(pgPool)

	argsWithProg := os.Args

	if len(argsWithProg) < 1 {
		return false
	}

	if argsWithProg[1] == "create-super-user" {
		if len(argsWithProg) < 4 { //nolint
			slog.Error("использование: create-super-user [login] [password]")

			return true
		}

		err := userDB.CreateUser(
			database.User{
				Login:       argsWithProg[2],
				Name:        argsWithProg[2],
				Permissions: []string{database.PermSuperUser},
			},
			argsWithProg[3],
		)
		if err != nil {
			panic("failed to create super user: " + err.Error())
		}

		slog.Warn("Создан супер пользователь, убедитесь что удалите его после использования")

		return true
	}

	return true
}
