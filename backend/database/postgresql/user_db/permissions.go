package user_db

import (
	"context"
	"fmt"

	"github.com/univers106/ITI/database"
	"github.com/univers106/ITI/database/postgresql"
)

func (db *UserDatabase) UserAddPermission(login string, permission string) error {
	query := `
		UPDATE public.users
		SET permissions = array_append(permissions, $2)
		WHERE login = $1 and not $2 = any(permissions);
	`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	res, err := db.pool.Exec(ctx, query, login, permission)
	if err != nil {
		return fmt.Errorf("failed to add permission: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("%w; or permission already exists", database.ErrUserNotFound)
	}

	return nil
}

func (db *UserDatabase) UserRemovePermission(login string, permission string) error {
	query := `
		UPDATE public.users
		SET permissions = array_remove(permissions, $2)
		WHERE login = $1 and $2 = any(permissions);
	`

	ctx, cancel := context.WithTimeout(context.Background(), postgresql.ReqTimeout)
	defer cancel()

	res, err := db.pool.Exec(ctx, query, login, permission)
	if err != nil {
		return fmt.Errorf("failed to remove permission: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("%w; or no user found", database.ErrPermissionNotFound)
	}

	return nil
}
