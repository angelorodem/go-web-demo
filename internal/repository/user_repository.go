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

type UserRepositoryInterface interface {
	CreateUser(usr *domain.User) error
	DeleteUser(email string) error
	ReadUser(email string) (*domain.User, error)
	UpdateUsername(email string, new_username string) error
}

// UserRepository handles all database operations for users
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(usr *domain.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx,
		"INSERT INTO users (username, email, password_hash) values (?, ?, ?)", usr.Username, usr.Email, usr.Password_hash)

	return err
}

func (r *UserRepository) DeleteUser(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE email == ?", email)

	return err
}

func (r *UserRepository) ReadUser(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	row := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE email == ?", email)

	var u domain.User

	if err := row.Scan(&u.Id, &u.Email, &u.Username, &u.Password_hash); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) UpdateUsername(email string, new_username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	res, err := r.db.ExecContext(ctx, "UPDATE users SET username = ? WHERE email == ?", new_username, email)

	if err != nil {
		return err
	} else if n, err := res.RowsAffected(); err != nil || n <= 0 {
		return fmt.Errorf("nothing has changed (no user)")
	}

	return nil
}
