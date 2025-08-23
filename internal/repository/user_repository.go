// Create user, hash password
// Login user, check password agains hash, if its true return "FAKE_AUTH_TOKEN"
// Delete user
// No real token is used/generated just a fixed string.
package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"web/example/internal/domain"
)

func CreateUser(db *sql.DB, usr *domain.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx,
		"INSERT INTO users (username, email, password_hash) values (?, ?, ?)", usr.Username, usr.Email, usr.Password_hash)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(db *sql.DB, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	_, err := db.ExecContext(ctx, "DELETE FROM users WHERE email == ?", email)

	if err != nil {
		return err
	}

	return nil
}

func GetUser(db *sql.DB, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	row := db.QueryRowContext(ctx, "SELECT * FROM users WHERE email == ?", email)

	var u domain.User

	var id int
	if err := row.Scan(&id, &u.Email, &u.Username, &u.Password_hash); err != nil {
		return nil, err
	}
	return &u, nil
}

func UpdateUsername(db *sql.DB, email string, new_username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	res, err := db.ExecContext(ctx, "UPDATE users SET username = ? WHERE email == ?", new_username, email)

	if err != nil {
		return err
	} else if n, err := res.RowsAffected(); err != nil || n <= 0 {
		return fmt.Errorf("nothing has changed (no user)")
	}

	return nil
}
